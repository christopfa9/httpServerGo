package handler

import (
	"net"
	"worker/internal/status"
	"worker/internal/utils"
)

func HandleStatus(params map[string]string, conn net.Conn) {
	data, err := status.Marshal()
	if err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "application/json", string(data))
}
