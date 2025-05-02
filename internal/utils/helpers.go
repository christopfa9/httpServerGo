package utils

// TODO: Implement helpers.go (Shared utility functions)
//
// [ ] Import necessary packages:
//     - fmt, strings, net, encoding/json, crypto/sha256, encoding/hex, etc.
//
// [ ] Define helper functions such as:
//
//     [ ] ParseQueryParams(rawQuery string) map[string]string
//         - Parses URL query string into key-value map
//
//     [ ] WriteHTTPResponse(conn net.Conn, statusCode int, contentType string, body string)
//         - Builds and writes a basic HTTP/1.0 response
//
//     [ ] SanitizeFileName(name string) string
//         - Prevents path traversal (removes "../", etc.)
//
//     [ ] JSONResponse(data any) (string, error)
//         - Converts a struct or map to a pretty-printed JSON string
//
//     [ ] SHA256Hash(input string) string
//         - Computes SHA-256 hash and returns it as hex string
//
// [ ] Ensure proper error handling and logging where needed

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"path/filepath"
	"strings"
)

// ParseQueryParams convierte una cadena "a=1&b=foo" en un map[string]string.
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

// WriteHTTPResponse genera y envía una respuesta HTTP/1.0 básica.
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

// SanitizeFileName devuelve solo el nombre base del archivo, eliminando cualquier ruta.
func SanitizeFileName(name string) string {
	// filepath.Base descarta cualquier prefijo de directorio
	base := filepath.Base(name)
	// opcional: eliminar caracteres sospechosos
	return strings.ReplaceAll(base, string(filepath.Separator), "")
}

// JSONResponse serializa data (struct, map, slice, etc.) a JSON indentado.
// Si data es ya un []byte (p.ej. el JSON crudo de status.Marshal), lo devuelve sin tocar.
func JSONResponse(data any) (string, error) {
	if raw, ok := data.([]byte); ok {
		// ya está listo para enviar
		return string(raw), nil
	}
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// SHA256Hash calcula el hash SHA-256 de input y lo devuelve en hex.
func SHA256Hash(input string) string {
	sum := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sum[:])
}
