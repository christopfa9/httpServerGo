package server

import (
	"bufio"
	"net"
	"strings"
	"time"

	"worker/internal/status"
	"worker/internal/utils"
	"worker/internal/handler"
)

// HandleConnection processes a simple HTTP/1.0 request, delegates the execution
// of each command to its worker pool, and responds using utility functions from utils.
func HandleConnection(conn net.Conn) {
	status.IncTotalConnections()
	status.IncActiveHandlers()
	defer status.DecActiveHandlers()

	// Timeout to avoid blocking indefinitely
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	reader := bufio.NewReader(conn)

	// 1) Read and validate request line
	line, err := reader.ReadString('\n')
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "400 Bad Request\n")
		return
	}
	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) < 3 {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "400 Bad Request\n")
		return
	}
	if parts[0] != "GET" {
		utils.WriteHTTPResponse(conn, 405, "text/plain", "405 Method Not Allowed\n")
		return
	}

	// 2) Extract path and query parameters
	rawURI := parts[1]
	path, query := rawURI, ""
	if idx := strings.Index(rawURI, "?"); idx != -1 {
		path = rawURI[:idx]
		query = rawURI[idx+1:]
	}
	params := utils.ParseQueryParams(query)

	// 3) Dispatch based on route, enqueueing to worker pools when applicable
	var (
		payload interface{}
		cmdErr  error
	)

	switch path {
	case "/ping":
		// Healthcheck: respond immediately
		utils.WriteHTTPResponse(conn, 200, "text/plain", "pong\n")
		return

	case "/fibonacci":
		handler.HandleFibonacci(params, conn)
		return

	case "/createfile":
		handler.HandleCreateFile(params, conn)
		return

	case "/deletefile":
		handler.HandleDeleteFile(params, conn)
		return

	case "/reverse":
		handler.HandleReverse(params, conn)
		return

	case "/toupper":
		handler.HandleToUpper(params, conn)
		return

	case "/random":
		handler.HandleRandom(params, conn)
		return

	case "/timestamp":
		handler.HandleTimestamp(params, conn)
		return

	case "/hash":
		handler.HandleHash(params, conn)
		return

	case "/simulate":
		handler.HandleSimulate(params, conn)
		return

	case "/sleep":
		handler.HandleSleep(params, conn)
		return

	case "/loadtest":
		handler.HandleLoadTest(params, conn)
		return	

	case "/computepi":
		handler.HandleComputePi(params, conn)
		return

	case "/pow":
		handler.HandlePow(params, conn)
		return

	case "/status":
		handler.HandleStatus(params, conn)
		return

	case "/help":
		handler.HandleHelp(params, conn)
		return

	default:
		utils.WriteHTTPResponse(conn, 404, "text/plain", "404 Not Found\n")
		return
	}

	// 4) Handle command errors
	if cmdErr != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}

	// 5) Serialize payload and respond
	contentType := "text/plain"
	var bodyStr string
	if s, ok := payload.(string); ok {
		bodyStr = s
	} else {
		contentType = "application/json"
		bodyStr, _ = utils.JSONResponse(payload)
	}

	utils.WriteHTTPResponse(conn, 200, contentType, bodyStr)
}
