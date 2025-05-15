package commands

import (
	"time"
)

// Timestamp devuelve la hora actual del sistema en formato ISO-8601 (RFC3339).
func Timestamp() (string, error) {
	now := time.Now().UTC()
	// time.RFC3339 da un string como "2006-01-02T15:04:05Z07:00"
	return now.Format(time.RFC3339), nil
}
