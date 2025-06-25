package handler

import (
	"net"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleTimestamp(params map[string]string, conn net.Conn) {
	respCh := make(chan pool.TimestampResult)
	pool.TimestampJobs <- pool.TimestampJob{Resp: respCh}

	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
