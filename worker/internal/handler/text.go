package handler

import (
	"net"
	"worker/internal/pool"
	"worker/internal/utils"
)

func HandleReverse(params map[string]string, conn net.Conn) {
	text := params["text"]

	respCh := make(chan pool.ReverseResult)
	pool.ReverseJobs <- pool.ReverseJob{
		Text: text,
		Resp: respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}

func HandleToUpper(params map[string]string, conn net.Conn) {
	text := params["text"]

	respCh := make(chan pool.ToUpperResult)
	pool.ToUpperJobs <- pool.ToUpperJob{
		Text: text,
		Resp: respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
