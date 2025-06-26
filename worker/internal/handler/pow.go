package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
	"worker/internal/status"
)

func HandlePow(params map[string]string, conn net.Conn) {
	prefix := params["prefix"]
	maxTrials, err := strconv.Atoi(params["maxTrials"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'maxTrials' parameter\n")
		return
	}

	respCh := make(chan pool.PowResult)
	pool.PowJobs <- pool.PowJob{
		Prefix:    prefix,
		MaxTrials: maxTrials,
		Resp:      respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	status.IncCompletedTasks()
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
