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
	// Lista de direcciones “host:port” de los Workers
	workers    []string
	nextIndex  int
	workersMtx sync.Mutex

	// Intervalo para hacer healthchecks
	healthcheckInterval = 10 * time.Second
)

// -----------------------------
//
//  1. ParseWorkersEnv convierte la cadena “host1:8080,host2:8080,...”
//     en un slice de strings limpiados.
//
//     Ejemplo de ENV: WORKERS="127.0.0.1:8080,127.0.0.1:8081"
//
// -----------------------------
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

// -----------------------------
//  2. InitDispatcher inicializa la lista de Workers y arranca
//     la rutina de healthchecks periódicos.
//
// -----------------------------
func InitDispatcher(list []string) {
	workers = make([]string, len(list))
	copy(workers, list)
	nextIndex = 0

	// Arrancar goroutine para healthchecks
	go healthcheckLoop()
}

// healthcheckLoop consulta periódicamente el endpoint /ping de cada Worker.
// Si un Worker no responde, lo omite temporalmente (lo saca de la rotación).
func healthcheckLoop() {
	ticker := time.NewTicker(healthcheckInterval)
	defer ticker.Stop()

	for range ticker.C {
		workersMtx.Lock()
		var alive []string
		for _, w := range workers {
			conn, err := net.DialTimeout("tcp", w, 1*time.Second)
			if err != nil {
				// Este Worker no está disponible
				continue
			}
			// Enviamos un “GET /ping HTTP/1.0\r\n\r\n”
			fmt.Fprintf(conn, "GET /ping HTTP/1.0\r\nHost: %s\r\n\r\n", w)
			// Leemos la primera línea de respuesta
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
		// Si hay al menos un Worker vivo, actualizamos la lista. Si no, mantenemos la anterior.
		if len(alive) > 0 {
			workers = alive
			if nextIndex >= len(workers) {
				nextIndex = 0
			}
		}
		workersMtx.Unlock()
	}
}

// -----------------------------
//  3. pickNextWorker devuelve la dirección “host:port” del Worker
//     según round-robin (y avanza el índice).
//     Si no hay Workers listos, retorna cadena vacía.
//
// -----------------------------
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

// -----------------------------
//  4. HandleConnection recibe cada conexión del cliente, lee la petición
//     HTTP/1.0 (únicamente “GET /ruta?...”), elige un Worker, le reenvía
//     íntegra la petición, lee la respuesta completa y la retransmite al
//     cliente original.
//
// -----------------------------
func HandleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	// 4.1) Leemos la primera línea “GET /algo?x=1 HTTP/1.0”
	clientReader := bufio.NewReader(clientConn)
	requestLine, err := clientReader.ReadString('\n')
	if err != nil {
		return
	}
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) < 3 || parts[0] != "GET" {
		// Solo aceptamos GET
		fmt.Fprint(clientConn, "HTTP/1.0 405 Method Not Allowed\r\n\r\n")
		return
	}
	uri := parts[1]

	// 4.2) Recolectar las demás cabeceras (hasta línea vacía)
	var headers []string
	for {
		line, err := clientReader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.TrimSpace(line) == "" {
			// fin de cabeceras
			break
		}
		headers = append(headers, line)
	}

	// 4.3) Elegimos el siguiente Worker disponible (round-robin)
	workerAddr := pickNextWorker()
	if workerAddr == "" {
		// No hay ningún Worker vivo
		fmt.Fprint(clientConn, "HTTP/1.0 503 Service Unavailable\r\n\r\n")
		return
	}

	// 4.4) Abrimos conexión al Worker
	workerConn, err := net.Dial("tcp", workerAddr)
	if err != nil {
		// Falló conexión al Worker
		fmt.Fprint(clientConn, "HTTP/1.0 502 Bad Gateway\r\n\r\n")
		return
	}
	defer workerConn.Close()

	// 4.5) Reenviamos línea de petición al Worker (manteniendo HTTP/1.0)
	fmt.Fprintf(workerConn, "GET %s HTTP/1.0\r\n", uri)
	// Reenviamos el resto de cabeceras
	for _, h := range headers {
		fmt.Fprint(workerConn, h)
	}
	// Línea en blanco para terminar request
	fmt.Fprint(workerConn, "\r\n")

	// 4.6) Leemos ENTENDIDA la respuesta completa del Worker y la reenviamos al cliente
	//       Copiamos byte a byte hasta EOF
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
