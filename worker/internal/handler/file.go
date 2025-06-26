package handler

import (
	"net"
	"strconv"
	"worker/internal/pool"
	"worker/internal/utils"
	"worker/internal/status"
)

func HandleCreateFile(params map[string]string, conn net.Conn) {
	name := params["name"]
	content := params["content"]
	repeat, _ := strconv.Atoi(params["repeat"])

	respCh := make(chan pool.CreateFileResult)
	pool.CreateFileJobs <- pool.CreateFileJob{
		Name:    name,
		Content: content,
		Repeat:  repeat,
		Resp:    respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	status.IncCompletedTasks()
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}

func HandleDeleteFile(params map[string]string, conn net.Conn) {
	name := params["name"]

	respCh := make(chan pool.DeleteFileResult)
	pool.DeleteFileJobs <- pool.DeleteFileJob{
		Name: name,
		Resp: respCh,
	}
	res := <-respCh
	if res.Err != nil {
		utils.WriteHTTPResponse(conn, 500, "text/plain", "500 Internal Server Error\n")
		return
	}
	utils.WriteHTTPResponse(conn, 200, "text/plain", res.Value)
}
