package commands

// TODO: Implement hash.go (Handles /hash?text=)
//
// [ ] Import necessary packages:
//     - fmt, net, crypto/sha256, encoding/hex, strings
//
// [ ] Define function HandleHash(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "text" parameter:
//     - Required
//     - Return 400 Bad Request if missing
//
// [ ] Compute SHA-256 hash of the input text
//
// [ ] Encode the hash in hexadecimal format
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: SHA-256 hash in hex
//
// [ ] Handle missing parameters and internal errors
//
// [ ] Optionally log the request and result

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hash calcula el SHA-256 del texto de entrada y devuelve su representación hexadecimal.
// Retorna error solo si ocurre un fallo inesperado (muy poco probable).
func Hash(text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("el parámetro 'text' es obligatorio")
	}
	// 1) Obtener bytes del texto y calcular SHA-256
	hashBytes := sha256.Sum256([]byte(text))
	// 2) Codificar en hexadecimal
	hexHash := hex.EncodeToString(hashBytes[:])
	return hexHash, nil
}
