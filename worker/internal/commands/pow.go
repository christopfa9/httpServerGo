package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// Pow busca un entero x en [0, maxTrials) tal que el SHA-256 de su representación
// como string comience con el prefijo hexadecimal dado.
//   - prefix: debe ser una cadena hex válida (sin “0x”).
//   - maxTrials: límite de intentos (debe ser > 0).
//
// Retorna el primer x que cumpla la condición y su hash, o error si no se encuentra.
func Pow(prefix string, maxTrials int) (string, error) {
	// Validar que prefix sea hex válido
	if _, err := hex.DecodeString(prefix); err != nil {
		return "", fmt.Errorf("prefix no es un hex válido: %v", err)
	}
	if maxTrials <= 0 {
		return "", fmt.Errorf("maxTrials debe ser > 0, recibí %d", maxTrials)
	}

	for i := 0; i < maxTrials; i++ {
		// Convertimos i a string para hashear
		val := strconv.Itoa(i)
		sum := sha256.Sum256([]byte(val))
		hexsum := hex.EncodeToString(sum[:])
		if len(hexsum) >= len(prefix) && hexsum[:len(prefix)] == prefix {
			// Devolvemos el número encontrado y su hash
			return fmt.Sprintf("found=%d hash=%s", i, hexsum), nil
		}
	}

	return "", fmt.Errorf("no se halló ningún valor en %d intentos", maxTrials)
}
