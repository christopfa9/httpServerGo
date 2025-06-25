package internal

import (
	"bufio"
	"fmt"
	"net"
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

	// 4.1) Read the first line: “GET /something?x=1 HTTP/1.0”
	clientReader := bufio.NewReader(clientConn)
	requestLine, err := clientReader.ReadString('\n')
	if err != nil {
		return
	}
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) < 3 || parts[0] != "GET" {
		// Only GET is allowed
		fmt.Fprint(clientConn, "HTTP/1.0 405 Method Not Allowed\r\n\r\n")
		return
	}
	uri := parts[1]

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
