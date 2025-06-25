package commands

import (
	"fmt"
)

// Fibonacci computes the n-th number in the Fibonacci sequence and
// returns the result as a string. Returns an error if n < 0.
func Fibonacci(n int) (string, error) {
	if n < 0 {
		return "", fmt.Errorf("the 'num' parameter must be >= 0, received %d", n)
	}
	result := fib(n)
	// Optional: log.Printf("Fibonacci(%d) = %d", n, result)
	return fmt.Sprintf("%d", result), nil
}

// fib is the pure recursive implementation of Fibonacci.
// For large n, this grows exponentially; for n > 40,
// consider memoization or an iterative approach.
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
