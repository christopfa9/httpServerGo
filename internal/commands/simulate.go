package commands

import (
	"fmt"
	"time"
)

// Simulate “ejecuta” una tarea durmiendo durante los segundos especificados.
// El parámetro task es opcional y se incluye en el mensaje de confirmación.
func Simulate(seconds int, task string) (string, error) {
	if seconds < 0 {
		return "", fmt.Errorf("el parámetro 'seconds' debe ser >= 0, recibí %d", seconds)
	}

	// Simulación: dormimos la cantidad de segundos indicada
	time.Sleep(time.Duration(seconds) * time.Second)

	// Construimos el mensaje de confirmación
	desc := "tarea"
	if task != "" {
		desc = task
	}
	msg := fmt.Sprintf("Simulated task: %s for %d seconds", desc, seconds)
	return msg, nil
}
