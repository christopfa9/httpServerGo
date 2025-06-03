package commands

import (
	"fmt"
	"math/rand"
	"time"
)

// Random genera un slice de 'count' números enteros aleatorios en el rango [min, max].
// Retorna un error si los parámetros son inválidos.
func Random(count, min, max int) ([]int, error) {
	// Validación de parámetros
	if count <= 0 {
		return nil, fmt.Errorf("el parámetro 'count' debe ser > 0, recibí %d", count)
	}
	if min > max {
		return nil, fmt.Errorf("el parámetro 'max' (%d) debe ser >= 'min' (%d)", max, min)
	}

	// Sembramos el generador; para un servidor real quizá quieras hacerlo una sola vez en init().
	rand.Seed(time.Now().UnixNano())

	// Generamos los números aleatorios
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}
	return nums, nil
}
