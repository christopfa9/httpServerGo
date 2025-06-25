package commands

import (
	"fmt"
	"time"
)

// Sleep pauses execution for the specified number of seconds.
// Returns a confirmation message or an error if the parameter is invalid.
func Sleep(seconds int) (string, error) {
	if seconds < 0 {
		return "", fmt.Errorf("the 'seconds' parameter must be >= 0, received %d", seconds)
	}
	// Perform the pause
	time.Sleep(time.Duration(seconds) * time.Second)
	// Prepare confirmation message
	msg := fmt.Sprintf("Slept for %d seconds", seconds)
	return msg, nil
}
