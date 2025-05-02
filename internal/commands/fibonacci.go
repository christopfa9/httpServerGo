package commands

// TODO: Implement fibonacci.go (Handles /fibonacci?num=N)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, strings
//
// [ ] Define function HandleFibonacci(conn net.Conn, params map[string]string)
//
// [ ] Extract "num" parameter from params
//     - Validate that it exists and is a valid integer
//     - Return 400 Bad Request if invalid or missing
//
// [ ] Implement a recursive function fibonacci(n int) int
//     - Optionally use memoization if performance is a concern
//
// [ ] Compute the Fibonacci number for n
//
// [ ] Write the HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: result as string
//
// [ ] Handle errors gracefully and return appropriate HTTP error codes
//
// [ ] Log request and result (optional)

import (
	"fmt"
)

// Fibonacci calcula el n-ésimo número de la serie de Fibonacci y
// devuelve el resultado como string. Retorna error si n < 0.
func Fibonacci(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("el parámetro 'num' debe ser >= 0, recibí %d", n)
	}
	resultado := fib(n)
	// Optional: log.Printf("Fibonacci(%d) = %d", n, resultado)
	return fmt.Sprintf("%d", resultado), nil
}

// fib es la implementación recursiva pura de Fibonacci.
// Para n muy grandes esto crecerá exponencialmente; si lo pruebas
// con n > 40, quizá quieras memoización o un enfoque iterativo.
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
