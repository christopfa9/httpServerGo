package commands

import (
	"strings"
)

// ToUpper converts the input text string to uppercase and returns it.
// It does not return an error since any text (even empty) is valid.
func ToUpper(text string) (string, error) {
	return strings.ToUpper(text), nil
}
