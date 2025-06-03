package commands

import (
	"fmt"
	mathrand "math/rand"
	"time"
)

// ComputePi calcula una aproximación de π usando el método de Monte Carlo.
//   - `iters` es el número de puntos aleatorios a generar (debe ser > 0).
//
// Retorna la aproximación de π como string o un error si `iters` es inválido.
func ComputePi(iters int) (string, error) {
	if iters <= 0 {
		return "", fmt.Errorf("iters debe ser > 0, recibí %d", iters)
	}

	// Sembramos el generador con la hora actual para obtener diferentes secuencias.
	mathrand.Seed(time.Now().UnixNano())

	count := 0
	for i := 0; i < iters; i++ {
		x := mathrand.Float64() // valor aleatorio en [0, 1)
		y := mathrand.Float64()
		if x*x+y*y <= 1.0 {
			count++
		}
	}

	pi := 4.0 * float64(count) / float64(iters)
	// Formateamos el resultado como string con precisión estándar.
	return fmt.Sprintf("%f", pi), nil
}
