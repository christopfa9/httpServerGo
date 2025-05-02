package commands

// TODO: Implement deletefile.go (Handles /deletefile?name=)
//
// [ ] Import necessary packages:
//     - fmt, net, os, strings
//
// [ ] Define function HandleDeleteFile(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "name" parameter:
//     - Required
//     - Sanitize to prevent path traversal (no "../" etc.)
//
// [ ] Check if the file exists
//     - If not, respond with 404 Not Found
//
// [ ] Attempt to delete the file using os.Remove
//
// [ ] Write appropriate HTTP response:
//     - 200 OK if successfully deleted
//     - 500 Internal Server Error if deletion fails
//
// [ ] Handle missing or invalid parameters with 400 Bad Request
//
// [ ] Log result (optional)

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DeleteFile elimina el archivo indicado. Retorna un mensaje de confirmación o error.
func DeleteFile(name string) (string, error) {
	// 1) Validar y sanitizar filename
	if name == "" {
		return "", fmt.Errorf("el parámetro 'name' es obligatorio")
	}
	// Prevenir directory traversal: no permitimos separadores ni rutas relativas
	if strings.ContainsAny(name, `/\`) || filepath.Base(name) != name {
		return "", fmt.Errorf("nombre de archivo inválido: %q", name)
	}

	// 2) Comprobar existencia
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("archivo no encontrado: %q", name)
		}
		return "", fmt.Errorf("error al acceder al archivo %q: %w", name, err)
	}

	// 3) Intentar eliminar
	if err := os.Remove(name); err != nil {
		return "", fmt.Errorf("error al eliminar el archivo %q: %w", name, err)
	}

	// 4) Confirmación
	msg := fmt.Sprintf("Archivo %q eliminado con éxito", name)
	return msg, nil
}
