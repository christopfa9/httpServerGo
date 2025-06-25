package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type matmulRequest struct {
	A        [][]float64 `json:"A"`
	B        [][]float64 `json:"B"`
	RowStart int         `json:"row_start"`
	RowEnd   int         `json:"row_end"`
}

type matmulResponse struct {
	Rows [][]float64 `json:"rows"`
}

func HandleMatMul(conn net.Conn) {
	defer conn.Close()
	// Leer body JSON
	body, err := io.ReadAll(conn)
	if err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nError reading body\n")
		return
	}
	var req matmulRequest
	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid JSON\n")
		return
	}
	// Validar dimensiones
	if len(req.A) == 0 || len(req.B) == 0 || len(req.A[0]) != len(req.B) {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nDimension mismatch\n")
		return
	}
	m := len(req.A)
	n := len(req.B)
	p := len(req.B[0])
	if req.RowStart < 0 || req.RowEnd > m || req.RowStart >= req.RowEnd {
		fmt.Fprintf(conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/plain\r\n\r\nInvalid row range\n")
		return
	}
	// Calcular solo las filas asignadas
	rows := make([][]float64, req.RowEnd-req.RowStart)
	for i := req.RowStart; i < req.RowEnd; i++ {
		row := make([]float64, p)
		for j := 0; j < p; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += req.A[i][k] * req.B[k][j]
			}
			row[j] = sum
		}
		rows[i-req.RowStart] = row
	}
	resp := matmulResponse{Rows: rows}
	data, _ := json.Marshal(resp)
	fmt.Fprintf(conn, "HTTP/1.0 200 OK\r\nContent-Type: application/json\r\n\r\n%s", data)
}
