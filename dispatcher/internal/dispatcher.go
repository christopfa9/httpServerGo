package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	// List of “host:port” addresses of the Workers
	workers    []string
	nextIndex  int
	workersMtx sync.Mutex

	// Interval for performing health checks
	healthcheckInterval = 10 * time.Second
)

//
//  1. ParseWorkersEnv converts the string “host1:8080,host2:8080,...”
//     into a cleaned slice of strings.
//
//     Example ENV: WORKERS="127.0.0.1:8080,127.0.0.1:8081"
//
func ParseWorkersEnv(env string) []string {
	parts := strings.Split(env, ",")
	var result []string
	for _, part := range parts {
		if p := strings.TrimSpace(part); p != "" {
			result = append(result, p)
		}
	}
	return result
}

//
//  2. InitDispatcher initializes the list of Workers and starts
//     the periodic health check routine.
//
func InitDispatcher(list []string) {
	workers = make([]string, len(list))
	copy(workers, list)
	nextIndex = 0

	// Start goroutine for health checks
	go healthcheckLoop()
}

// healthcheckLoop periodically checks the /ping endpoint of each Worker.
// If a Worker doesn't respond, it's temporarily removed from rotation.
func healthcheckLoop() {
	ticker := time.NewTicker(healthcheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		workersMtx.Lock()
		var alive []string
		for _, w := range workers {
			conn, err := net.DialTimeout("tcp", w, 1*time.Second)
			if err != nil {
				// This Worker is not available
				continue
			}
			// Send “GET /ping HTTP/1.0\r\n\r\n”
			fmt.Fprintf(conn, "GET /ping HTTP/1.0\r\nHost: %s\r\n\r\n", w)
			reader := bufio.NewReader(conn)
			line, err := reader.ReadString('\n')
			conn.Close()
			if err != nil {
				continue
			}
			if strings.Contains(line, "200") {
				alive = append(alive, w)
			}
		}
		// If at least one Worker is alive, update the list. Otherwise, keep the current one.
		if len(alive) > 0 {
			workers = alive
			if nextIndex >= len(workers) {
				nextIndex = 0
			}
		}
		workersMtx.Unlock()
	}
}

//
//  3. pickNextWorker returns the “host:port” of the next Worker
//     using round-robin (and advances the index).
//     If there are no ready Workers, returns an empty string.
//
func pickNextWorker() string {
	workersMtx.Lock()
	defer workersMtx.Unlock()

	if len(workers) == 0 {
		return ""
	}
	w := workers[nextIndex]
	nextIndex = (nextIndex + 1) % len(workers)
	return w
}

//
//  4. HandleConnection receives a client connection, reads the
//     HTTP/1.0 request (“GET /route?...”), picks a Worker, forwards
//     the entire request, reads the full response, and relays it
//     back to the original client.
//
func HandleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	clientReader := bufio.NewReader(clientConn)
	requestLine, err := clientReader.ReadString('\n')
	if err != nil {
		return
	}
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) < 3 || parts[0] != "GET" {
		fmt.Fprint(clientConn, "HTTP/1.0 405 Method Not Allowed\r\n\r\n")
		return
	}
	uri := parts[1]

	// Detectar si es /computepi
	if strings.HasPrefix(uri, "/computepi") {
		distributedComputePi(clientConn, uri, clientReader)
		return
	}

	// Detectar si es /pow
	if strings.HasPrefix(uri, "/pow") {
		distributedPow(clientConn, uri, clientReader)
		return
	}

	// Endpoint especial para /workers: agrega antes del default
	if uri == "/workers" {
		reports, err := GetWorkersStatus()
		if err != nil {
			fmt.Fprintf(clientConn, "HTTP/1.0 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\nError: %v\n", err)
			return
		}
		jsonData, _ := json.MarshalIndent(reports, "", "  ")
		fmt.Fprintf(clientConn, "HTTP/1.0 200 OK\r\nContent-Type: application/json\r\n\r\n%s", string(jsonData))
		return
	}

	// 4.2) Collect the remaining headers (until empty line)
	var headers []string
	for {
		line, err := clientReader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.TrimSpace(line) == "" {
			// end of headers
			break
		}
		headers = append(headers, line)
	}

	// 4.3) Select the next available Worker (round-robin)
	workerAddr := pickNextWorker()
	if workerAddr == "" {
		// No live Worker available
		fmt.Fprint(clientConn, "HTTP/1.0 503 Service Unavailable\r\n\r\n")
		return
	}

	// 4.4) Open connection to the Worker
	workerConn, err := net.Dial("tcp", workerAddr)
	if err != nil {
		// Failed to connect to the Worker
		fmt.Fprint(clientConn, "HTTP/1.0 502 Bad Gateway\r\n\r\n")
		return
	}
	defer workerConn.Close()

	// 4.5) Forward request line to Worker (keeping HTTP/1.0)
	fmt.Fprintf(workerConn, "GET %s HTTP/1.0\r\n", uri)
	// Forward the rest of the headers
	for _, h := range headers {
		fmt.Fprint(workerConn, h)
	}
	// Empty line to end the request
	fmt.Fprint(workerConn, "\r\n")

	// 4.6) Read full response from Worker and relay it to the client
	//      Copy byte by byte until EOF
	workerReader := bufio.NewReader(workerConn)
	for {
		buf := make([]byte, 4096)
		n, err := workerReader.Read(buf)
		if n > 0 {
			clientConn.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
}

// distributedComputePi divide la tarea entre los workers y ensambla el resultado
func distributedComputePi(clientConn net.Conn, uri string, clientReader *bufio.Reader) {
	workersMtx.Lock()
	activeWorkers := make([]string, len(workers))
	copy(activeWorkers, workers)
	workersMtx.Unlock()
	if len(activeWorkers) == 0 {
		fmt.Fprint(clientConn, "HTTP/1.0 503 Service Unavailable\r\n\r\n")
		return
	}

	u, _ := url.Parse(uri)
	params := u.Query()
	itersStr := params.Get("iters")
	iters, err := strconv.Atoi(itersStr)
	if err != nil || iters <= 0 {
		fmt.Fprint(clientConn, "HTTP/1.0 400 Bad Request\r\n\r\nInvalid iters parameter\n")
		return
	}

	// Dividir iteraciones
	perWorker := iters / len(activeWorkers)
	extra := iters % len(activeWorkers)

	type piResp struct {
		Value float64 `json:"value"`
		Err   string  `json:"err"`
	}
	results := make([]float64, len(activeWorkers))
	errCh := make(chan error, len(activeWorkers))
	respCh := make(chan struct{
		idx int
		val float64
	}, len(activeWorkers))

	for i, w := range activeWorkers {
		wIters := perWorker
		if i < extra {
			wIters++
		}
		go func(idx int, workerAddr string, n int) {
			// Construir sub-request
			wuri := fmt.Sprintf("/computepi?iters=%d", n)
			conn, err := net.DialTimeout("tcp", workerAddr, 2*time.Second)
			if err != nil {
				errCh <- err
				return
			}
			defer conn.Close()
			fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", wuri, workerAddr)
			reader := bufio.NewReader(conn)
			// Leer status
			line, err := reader.ReadString('\n')
			if err != nil || !strings.Contains(line, "200") {
				errCh <- fmt.Errorf("bad response from worker")
				return
			}
			// Leer headers hasta línea vacía
			for {
				h, _ := reader.ReadString('\n')
				if strings.TrimSpace(h) == "" {
					break
				}
			}
			// Leer body (asumimos float64 en texto plano)
			body, _ := reader.ReadString('\n')
			val, err := strconv.ParseFloat(strings.TrimSpace(body), 64)
			if err != nil {
				errCh <- err
				return
			}
			respCh <- struct{idx int; val float64}{idx, val}
		}(i, w, wIters)
	}

	total := 0.0
	done := 0
	for done < len(activeWorkers) {
		select {
		case r := <-respCh:
			results[r.idx] = r.val
			done++
		case <-time.After(3 * time.Second):
			errCh <- fmt.Errorf("timeout waiting for worker")
			done++
		}
	}
	for _, v := range results {
		total += v
	}
	// Promedio de los resultados parciales
	final := total / float64(len(activeWorkers))
	fmt.Fprintf(clientConn, "HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n%f\n", final)
}

// distributedPow divide la búsqueda de hash con prefijo entre los workers y responde con el primer resultado válido.
func distributedPow(clientConn net.Conn, uri string, clientReader *bufio.Reader) {
	workersMtx.Lock()
	activeWorkers := make([]string, len(workers))
	copy(activeWorkers, workers)
	workersMtx.Unlock()
	if len(activeWorkers) == 0 {
		fmt.Fprint(clientConn, "HTTP/1.0 503 Service Unavailable\r\n\r\n")
		return
	}

	u, _ := url.Parse(uri)
	params := u.Query()
	prefix := params.Get("prefix")
	maxTrialsStr := params.Get("maxTrials")
	maxTrials, err := strconv.Atoi(maxTrialsStr)
	if err != nil || maxTrials <= 0 {
		fmt.Fprint(clientConn, "HTTP/1.0 400 Bad Request\r\n\r\nInvalid maxTrials parameter\n")
		return
	}

	perWorker := maxTrials / len(activeWorkers)
	extra := maxTrials % len(activeWorkers)

	type powResp struct {
		idx  int
		val  string
		err  error
	}
	respCh := make(chan powResp, len(activeWorkers))

	for i, w := range activeWorkers {
		wTrials := perWorker
		if i < extra {
			wTrials++
		}
		go func(idx int, workerAddr string, trials int) {
			wuri := fmt.Sprintf("/pow?prefix=%s&maxTrials=%d", prefix, trials)
			conn, err := net.DialTimeout("tcp", workerAddr, 2*time.Second)
			if err != nil {
				respCh <- powResp{idx, "", err}
				return
			}
			defer conn.Close()
			fmt.Fprintf(conn, "GET %s HTTP/1.0\r\nHost: %s\r\n\r\n", wuri, workerAddr)
			reader := bufio.NewReader(conn)
			line, err := reader.ReadString('\n')
			if err != nil || !strings.Contains(line, "200") {
				respCh <- powResp{idx, "", fmt.Errorf("bad response from worker")}
				return
			}
			for {
				h, _ := reader.ReadString('\n')
				if strings.TrimSpace(h) == "" {
					break
				}
			}
			body, _ := reader.ReadString('\n')
			respCh <- powResp{idx, strings.TrimSpace(body), nil}
		}(i, w, wTrials)
	}

	var result string
	gotResult := false
	for i := 0; i < len(activeWorkers); i++ {
		resp := <-respCh
		if resp.err == nil && !gotResult {
			result = resp.val
			gotResult = true
		}
	}
	if gotResult {
		fmt.Fprintf(clientConn, "HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n%s\n", result)
	} else {
		fmt.Fprint(clientConn, "HTTP/1.0 404 Not Found\r\n\r\nNo valid hash found\n")
	}
}
