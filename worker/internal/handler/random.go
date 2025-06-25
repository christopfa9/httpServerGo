package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleRandom(params map[string]string, conn net.Conn) {
	count, e1 := strconv.Atoi(params["count"])
	min, e2 := strconv.Atoi(params["min"])
	max, e3 := strconv.Atoi(params["max"])

	if e1 != nil || e2 != nil || e3 != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid count/min/max parameters\n")
		return
	}

	respCh := make(chan pool.RandomResult)
	pool.RandomJobs <- pool.RandomJob{
		Count: count,
		Min:   min,
		Max:   max,
		Resp:  respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	json, _ := utils.JSONResponse(res.Value)
	utils.WriteHTTPResponse(conn, 200, "application/json", json)
}
