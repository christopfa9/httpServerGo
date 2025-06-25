package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleSimulate(params map[string]string, conn net.Conn) {
	seconds, err := strconv.Atoi(params["seconds"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'seconds' parameter\n")
		return
	}
	task := params["task"]

	respCh := make(chan pool.SimulateResult)
	pool.SimulateJobs <- pool.SimulateJob{
		Seconds: seconds,
		Task:    task,
		Resp:    respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}

func HandleSleep(params map[string]string, conn net.Conn) {
	seconds, err := strconv.Atoi(params["seconds"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'seconds' parameter\n")
		return
	}

	respCh := make(chan pool.SleepResult)
	pool.SleepJobs <- pool.SleepJob{
		Seconds: seconds,
		Resp:    respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
