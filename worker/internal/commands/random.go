package commands

import (
	"fmt"
	"math/rand"
	"time"
)

// Random generates a slice of 'count' random integers in the range [min, max].
// Returns an error if the parameters are invalid.
func Random(count, min, max int) ([]int, error) {
	// Parameter validation
	if count <= 0 {
		return nil, fmt.Errorf("the 'count' parameter must be > 0, received %d", count)
	}
	if min > max {
		return nil, fmt.Errorf("the 'max' parameter (%d) must be >= 'min' (%d)", max, min)
	}

	// Seed the random generator; for a real server you might want to do this once in init().
	rand.Seed(time.Now().UnixNano())

	// Generate the random numbers
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Intn(max-min+1) + min
	}
	return nums, nil
}
