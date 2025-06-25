package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Hash computes the SHA-256 hash of the input text and returns its hexadecimal representation.
// Returns an error only if an unexpected failure occurs (very unlikely).
func Hash(text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("the 'text' parameter is required")
	}
	// 1) Convert the text to bytes and compute SHA-256
	hashBytes := sha256.Sum256([]byte(text))
	// 2) Encode to hexadecimal
	hexHash := hex.EncodeToString(hashBytes[:])
	return hexHash, nil
}
