package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// Pow searches for an integer x in the range [0, maxTrials) such that the SHA-256 hash
// of its string representation starts with the given hexadecimal prefix.
//   - prefix: must be a valid hex string (without "0x").
//   - maxTrials: trial limit (must be > 0).
//
// Returns the first x that satisfies the condition and its hash, or an error if not found.
func Pow(prefix string, maxTrials int) (string, error) {
	// Validate that prefix is valid hex
	if _, err := hex.DecodeString(prefix); err != nil {
		return "", fmt.Errorf("prefix is not valid hex: %v", err)
	}
	if maxTrials <= 0 {
		return "", fmt.Errorf("maxTrials must be > 0, received %d", maxTrials)
	}

	for i := 0; i < maxTrials; i++ {
		// Convert i to string for hashing
		val := strconv.Itoa(i)
		sum := sha256.Sum256([]byte(val))
		hexsum := hex.EncodeToString(sum[:])
		if len(hexsum) >= len(prefix) && hexsum[:len(prefix)] == prefix {
			// Return the found number and its hash
			return fmt.Sprintf("found=%d hash=%s", i, hexsum), nil
		}
	}

	return "", fmt.Errorf("no value found within %d attempts", maxTrials)
}
