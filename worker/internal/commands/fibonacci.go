package commands

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
