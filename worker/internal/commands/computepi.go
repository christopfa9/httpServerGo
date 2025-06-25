package commands

import (
	"fmt"
	mathrand "math/rand"
	"time"
)

// ComputePi calculates an approximation of π using the Monte Carlo method.
//   - `iters` is the number of random points to generate (must be > 0).
//
// Returns the approximation of π as a string or an error if `iters` is invalid.
func ComputePi(iters int) (string, error) {
	if iters <= 0 {
		return "", fmt.Errorf("iters must be > 0, received %d", iters)
	}

	// Seed the random number generator with the current time to get different sequences.
	mathrand.Seed(time.Now().UnixNano())

	count := 0
	for i := 0; i < iters; i++ {
		x := mathrand.Float64() // random value in [0, 1)
		y := mathrand.Float64()
		if x*x+y*y <= 1.0 {
			count++
		}
	}

	pi := 4.0 * float64(count) / float64(iters)
	// Format the result as a string with standard precision.
	return fmt.Sprintf("%f", pi), nil
}
