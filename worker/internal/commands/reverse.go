package commands

// Reverse reverses the input text string and returns it.
// It does not return an error since any text (even empty) is valid.
func Reverse(text string) (string, error) {
	// Convert to a slice of runes to properly handle UTF-8 characters
	runes := []rune(text)
	// Swap elements from the ends toward the center
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
