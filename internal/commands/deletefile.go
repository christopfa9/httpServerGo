package commands

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
