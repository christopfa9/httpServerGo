package commands

import (
	"fmt"
	"sync"
	"time"
)

// LoadTest lanza 'tasks' goroutines concurrentes que duermen 'sleepSec' segundos cada una.
// Espera a que todas terminen antes de devolver la confirmación.
func LoadTest(tasks, sleepSec int) (string, error) {
	// 1) Validación de parámetros
	if tasks <= 0 {
		return "", fmt.Errorf("el parámetro 'tasks' debe ser > 0, recibí %d", tasks)
	}
	if sleepSec < 0 {
		return "", fmt.Errorf("el parámetro 'sleep' debe ser >= 0, recibí %d", sleepSec)
	}

	// 2) Lanzar goroutines concurrentes con WaitGroup
	var wg sync.WaitGroup
	wg.Add(tasks)
	for i := 0; i < tasks; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(sleepSec) * time.Second)
		}(i)
	}

	// 3) Esperar a que todas terminen
	wg.Wait()

	// 4) Construir mensaje de confirmación
	msg := fmt.Sprintf("Executed %d concurrent tasks sleeping %d seconds each", tasks, sleepSec)
	return msg, nil
}
