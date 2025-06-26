package handler

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"worker/internal/pool"
	"worker/internal/status"
)

func HandleComputePi(params map[string]string, conn net.Conn) {
	iters, err := strconv.Atoi(params["iters"])
	if err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid 'iters' parameter\n")
		return
	}
	log.Printf("Worker PID %d: iters=%d", os.Getpid(), iters)
	respCh := make(chan pool.ComputePiResult)
	pool.ComputePiJobs <- pool.ComputePiJob{
		Iters: iters,
		Resp:  respCh,
	}
	res := <-respCh
	if res.Err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 500 Internal Server Error\r\nContent-Type: text/plain\r\n\r\n500 Internal Server Error\n")
		return
	}
	// Incrementa el contador de tareas completadas
	status.IncCompletedTasks()
	// Responder solo el valor de pi en texto plano
	fmt.Fprintf(conn, "HTTP/1.0 200 OK\r\nContent-Type: text/plain\r\n\r\n%s\n", res.Value)
}
