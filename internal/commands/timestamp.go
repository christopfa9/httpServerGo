package commands

// TODO: Implement timestamp.go (Handles /timestamp)
//
// [ ] Import necessary packages:
//     - fmt, net, time
//
// [ ] Define function HandleTimestamp(conn net.Conn)
//
// [ ] Get current system time using time.Now()
//
// [ ] Format the time in ISO-8601 using time.Format(time.RFC3339)
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: formatted timestamp
//
// [ ] Handle and log errors (if any)
//
// [ ] Optionally log request and response

import (
	"time"
)

// Timestamp devuelve la hora actual del sistema en formato ISO-8601 (RFC3339).
func Timestamp() (string, error) {
	now := time.Now().UTC()
	// time.RFC3339 da un string como "2006-01-02T15:04:05Z07:00"
	return now.Format(time.RFC3339), nil
}
