package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DeleteFile deletes the specified file. Returns a confirmation message or an error.
func DeleteFile(name string) (string, error) {
	// 1) Validate and sanitize filename
	if name == "" {
		return "", fmt.Errorf("the 'name' parameter is required")
	}
	// Prevent directory traversal: do not allow path separators or relative paths
	if strings.ContainsAny(name, `/\`) || filepath.Base(name) != name {
		return "", fmt.Errorf("invalid filename: %q", name)
	}

	// 2) Check if file exists
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file not found: %q", name)
		}
		return "", fmt.Errorf("error accessing file %q: %w", name, err)
	}

	// 3) Attempt to delete
	if err := os.Remove(name); err != nil {
		return "", fmt.Errorf("error deleting file %q: %w", name, err)
	}

	// 4) Confirmation
	msg := fmt.Sprintf("File %q successfully deleted", name)
	return msg, nil
}
