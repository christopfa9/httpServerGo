package handler

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"worker/internal/pool"
	"worker/internal/status"
	"worker/internal/utils"
)

func HandlePow(params map[string]string, conn net.Conn) {
	prefix := params["prefix"]
	maxTrials, err := strconv.Atoi(params["maxTrials"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'maxTrials' parameter\n")
		return
	}

	fmt.Printf("[Worker PID %d] Ejecutando POW: prefix=%s, maxTrials=%d\n", os.Getpid(), prefix, maxTrials)

	respCh := make(chan pool.PowResult)
	pool.PowJobs <- pool.PowJob{
		Prefix:    prefix,
		MaxTrials: maxTrials,
		Resp:      respCh,
	}
	res := <-respCh
	if res.Err != nil {
		fmt.Printf("[Worker PID %d] POW falló: %v\n", os.Getpid(), res.Err)
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	fmt.Printf("[Worker PID %d] POW completado, resultado: %s\n", os.Getpid(), res.Value)
	status.IncCompletedTasks()
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
