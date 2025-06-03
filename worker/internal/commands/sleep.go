package commands

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
