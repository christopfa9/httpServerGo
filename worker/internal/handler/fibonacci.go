package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/status"
	"worker/internal/utils"
)

func HandleFibonacci(params map[string]string, conn net.Conn) {
	n, err := strconv.Atoi(params["num"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'num' parameter\n")
		return
	}

	respCh := make(chan pool.FibResult)
	pool.FibJobs <- pool.FibJob{
		N:    n,
		Resp: respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	// Incrementa el contador de tareas completadas
	status.IncCompletedTasks()
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
