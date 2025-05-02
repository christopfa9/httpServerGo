package commands

// TODO: Implement sleep.go (Handles /sleep?seconds=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, time
//
// [ ] Define function HandleSleep(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "seconds" parameter:
//     - Required, must be integer ≥ 0
//     - Return 400 Bad Request if missing or invalid
//
// [ ] Perform sleep using time.Sleep for the given duration
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: confirmation message (e.g., "Slept for X seconds")
//
// [ ] Handle input errors and respond appropriately
//
// [ ] Optionally log the sleep operation

import (
	"fmt"
	"time"
)

// Sleep detiene la ejecución durante el número de segundos especificado.
// Retorna un mensaje de confirmación o error si el parámetro es inválido.
func Sleep(seconds int) (string, error) {
	if seconds < 0 {
		return "", fmt.Errorf("el parámetro 'seconds' debe ser >= 0, recibí %d", seconds)
	}
	// Ejecutamos la pausa
	time.Sleep(time.Duration(seconds) * time.Second)
	// Preparamos mensaje de confirmación
	msg := fmt.Sprintf("Slept for %d seconds", seconds)
	return msg, nil
}
