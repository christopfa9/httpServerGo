package commands

import (
	"fmt"
	"time"
)

// Simulate "executes" a task by sleeping for the specified number of seconds.
// The 'task' parameter is optional and included in the confirmation message.
func Simulate(seconds int, task string) (string, error) {
	if seconds < 0 {
		return "", fmt.Errorf("the 'seconds' parameter must be >= 0, received %d", seconds)
	}

	// Simulation: sleep for the specified number of seconds
	time.Sleep(time.Duration(seconds) * time.Second)

	// Build the confirmation message
	desc := "task"
	if task != "" {
		desc = task
	}
	msg := fmt.Sprintf("Simulated task: %s for %d seconds", desc, seconds)
	return msg, nil
}
