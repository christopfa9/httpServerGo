package commands

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
