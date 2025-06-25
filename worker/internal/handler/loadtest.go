package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleLoadTest(params map[string]string, conn net.Conn) {
	tasks, err1 := strconv.Atoi(params["tasks"])
	sleep, err2 := strconv.Atoi(params["sleep"])

	if err1 != nil || err2 != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid tasks/sleep parameters\n")
		return
	}

	respCh := make(chan pool.LoadTestResult)
	pool.LoadTestJobs <- pool.LoadTestJob{
		Tasks:    tasks,
		SleepSec: sleep,
		Resp:     respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
