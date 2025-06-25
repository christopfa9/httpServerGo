package commands

import (
	"time"
)

// Timestamp returns the current system time in ISO-8601 (RFC3339) format.
func Timestamp() (string, error) {
	now := time.Now().UTC()
	// time.RFC3339 gives a string like "2006-01-02T15:04:05Z07:00"
	return now.Format(time.RFC3339), nil
}
