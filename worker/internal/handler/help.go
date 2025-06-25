package handler

import (
	"net"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleHelp(params map[string]string, conn net.Conn) {
	respCh := make(chan pool.HelpResult)
	pool.HelpJobs <- pool.HelpJob{Resp: respCh}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
