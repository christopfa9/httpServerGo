package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleComputePi(params map[string]string, conn net.Conn) {
	iters, err := strconv.Atoi(params["iters"])
	if err != nil {
		utils.WriteHTTPResponse(conn, 400, "text/plain", "Invalid 'iters' parameter\n")
		return
	}

	respCh := make(chan pool.ComputePiResult)
	pool.ComputePiJobs <- pool.ComputePiJob{
		Iters: iters,
		Resp:  respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
