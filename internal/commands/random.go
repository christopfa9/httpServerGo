package commands

// TODO: Implement random.go (Handles /random?count=&min=&max=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, math/rand, time, encoding/json
//
// [ ] Define function HandleRandom(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "count": required, integer > 0
//     - "min": required, integer
//     - "max": required, integer, must be >= min
//
// [ ] Seed the random number generator using time.Now().UnixNano()
//
// [ ] Generate a slice of count random integers in [min, max]
//
// [ ] Marshal the result as a JSON array
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: application/json
//     - Body: JSON array of random numbers
//
// [ ] Handle validation errors and respond with 400 Bad Request
//
// [ ] Handle internal errors with 500 Internal Server Error

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
