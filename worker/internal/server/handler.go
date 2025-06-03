package server

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"time"

	"worker/internal/status"
	"worker/internal/utils"
)

// HandleConnection procesa una petición HTTP/1.0 simple, delega la ejecución
// de cada comando a su pool de workers y responde usando los utilitarios de utils.
func HandleConnection(conn net.Conn) {
	status.IncTotalConnections()
	status.IncActiveHandlers()
	defer status.DecActiveHandlers()

	// Timeout para no bloquearse indefinidamente
	conn.SetDeadline(time.Now().Add(5 * time.Second))
	reader := bufio.NewReader(conn)

	// 1) Leer y validar línea de petición
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

	// 2) Extraer ruta y parámetros
	rawURI := parts[1]
	path, query := rawURI, ""
	if idx := strings.Index(rawURI, "?"); idx != -1 {
		path = rawURI[:idx]
		query = rawURI[idx+1:]
	}
	params := utils.ParseQueryParams(query)

	// 3) Dispatch según ruta, encolando en pools cuando corresponda
	var (
		payload interface{}
		cmdErr  error
	)

	switch path {
	case "/ping":
		// Healthcheck: responder inmediatamente
		utils.WriteHTTPResponse(conn, 200, "text/plain", "pong\n")
		return

	case "/fibonacci":
		n, err := strconv.Atoi(params["num"])
		if err != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'num' parameter\n")
			return
		}
		respCh := make(chan fibResult)
		fibJobs <- fibJob{n: n, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/createfile":
		name := params["name"]
		content := params["content"]
		repeat, _ := strconv.Atoi(params["repeat"])
		respCh := make(chan createFileResult)
		createFileJobs <- createFileJob{name: name, content: content, repeat: repeat, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/deletefile":
		name := params["name"]
		respCh := make(chan deleteFileResult)
		deleteFileJobs <- deleteFileJob{name: name, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/reverse":
		text := params["text"]
		respCh := make(chan reverseResult)
		reverseJobs <- reverseJob{text: text, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/toupper":
		text := params["text"]
		respCh := make(chan toUpperResult)
		toUpperJobs <- toUpperJob{text: text, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/random":
		cnt, e1 := strconv.Atoi(params["count"])
		mn, e2 := strconv.Atoi(params["min"])
		mx, e3 := strconv.Atoi(params["max"])
		if e1 != nil || e2 != nil || e3 != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid count/min/max parameters\n")
			return
		}
		respCh := make(chan randomResult)
		randomJobs <- randomJob{count: cnt, min: mn, max: mx, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/timestamp":
		respCh := make(chan timestampResult)
		timestampJobs <- timestampJob{resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/hash":
		text := params["text"]
		respCh := make(chan hashResult)
		hashJobs <- hashJob{text: text, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/simulate":
		secs, e := strconv.Atoi(params["seconds"])
		if e != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'seconds' parameter\n")
			return
		}
		task := params["task"]
		respCh := make(chan simulateResult)
		simulateJobs <- simulateJob{seconds: secs, task: task, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/sleep":
		secs, e := strconv.Atoi(params["seconds"])
		if e != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'seconds' parameter\n")
			return
		}
		respCh := make(chan sleepResult)
		sleepJobs <- sleepJob{seconds: secs, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/loadtest":
		tasks, e1 := strconv.Atoi(params["tasks"])
		sleepSec, e2 := strconv.Atoi(params["sleep"])
		if e1 != nil || e2 != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid tasks/sleep parameters\n")
			return
		}
		respCh := make(chan loadTestResult)
		loadTestJobs <- loadTestJob{tasks: tasks, sleepSec: sleepSec, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/computepi":
		iters, err := strconv.Atoi(params["iters"])
		if err != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'iters' parameter\n")
			return
		}
		respCh := make(chan computePiResult)
		computePiJobs <- computePiJob{iters: iters, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/pow":
		prefix := params["prefix"]
		maxTrials, err := strconv.Atoi(params["maxTrials"])
		if err != nil {
			utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'maxTrials' parameter\n")
			return
		}
		respCh := make(chan powResult)
		powJobs <- powJob{prefix: prefix, maxTrials: maxTrials, resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	case "/status":
		payload, cmdErr = status.Marshal()

	case "/help":
		respCh := make(chan helpResult)
		helpJobs <- helpJob{resp: respCh}
		res := <-respCh
		payload, cmdErr = res.value, res.err

	default:
		utils.WriteHTTPResponse(conn, 404, "text/plain", "404 Not Found\n")
		return
	}

	// 4) Manejar errores de comando
	if cmdErr != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}

	// 5) Serializar payload y responder
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
