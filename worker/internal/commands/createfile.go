package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateFile creates or truncates the specified file and writes the given content
// the specified number of times. Returns a confirmation message or an error.
func CreateFile(name, content string, repeat int) (string, error) {
	// 1) Validate and sanitize the filename
	if name == "" {
		return "", fmt.Errorf("the 'name' parameter is required")
	}
	// Prevent directory traversal: only allow names without path separators
	if strings.ContainsAny(name, `/\`) || filepath.Base(name) != name {
		return "", fmt.Errorf("invalid filename: %q", name)
	}

	// 2) Adjust repeat if it's invalid
	if repeat < 1 {
		repeat = 1
	}

	// 3) Create or truncate the file
	f, err := os.Create(name)
	if err != nil {
		return "", fmt.Errorf("error creating or truncating file %q: %w", name, err)
	}
	defer f.Close()

	// 4) Write the content 'repeat' times
	for i := 0; i < repeat; i++ {
		if _, err := f.WriteString(content); err != nil {
			return "", fmt.Errorf("error writing to file %q: %w", name, err)
		}
	}

	// 5) Confirmation
	msg := fmt.Sprintf("File %q successfully created/truncated (%d repetitions)", name, repeat)
	return msg, nil
}
