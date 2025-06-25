package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"path/filepath"
	"strings"
)

// ParseQueryParams converts a string like "a=1&b=foo" into a map[string]string.
func ParseQueryParams(rawQuery string) map[string]string {
	params := make(map[string]string)
	if rawQuery == "" {
		return params
	}
	pairs := strings.Split(rawQuery, "&")
	for _, p := range pairs {
		kv := strings.SplitN(p, "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}
	}
	return params
}

// WriteHTTPResponse generates and sends a basic HTTP/1.0 response.
func WriteHTTPResponse(conn net.Conn, statusCode int, contentType, body string) error {
	statusText := map[int]string{
		200: "OK",
		400: "Bad Request",
		404: "Not Found",
		405: "Method Not Allowed",
		500: "Internal Server Error",
	}
	reason, ok := statusText[statusCode]
	if !ok {
		reason = "Status"
	}
	header := fmt.Sprintf(
		"HTTP/1.0 %d %s\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n",
		statusCode, reason, contentType, len(body),
	)
	if _, err := conn.Write([]byte(header)); err != nil {
		return err
	}
	if _, err := conn.Write([]byte(body)); err != nil {
		return err
	}
	return nil
}

// SanitizeFileName returns only the base name of the file, removing any path prefix.
func SanitizeFileName(name string) string {
	// filepath.Base discards any directory prefix
	base := filepath.Base(name)
	// optionally: remove suspicious characters
	return strings.ReplaceAll(base, string(filepath.Separator), "")
}

// JSONResponse serializes data (struct, map, slice, etc.) to indented JSON.
// If data is already a []byte (e.g. raw JSON from status.Marshal), it returns it as is.
func JSONResponse(data any) (string, error) {
	if raw, ok := data.([]byte); ok {
		// already ready to send
		return string(raw), nil
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// SHA256Hash computes the SHA-256 hash of the input and returns it in hex format.
func SHA256Hash(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}
