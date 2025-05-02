package commands

// TODO: Implement createfile.go (Handles /createfile?name=&content=&repeat=)
//
// [ ] Import necessary packages:
//     - fmt, net, os, strconv, strings
//
// [ ] Define function HandleCreateFile(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "name": required, no path traversal (sanitize input)
//     - "content": required
//     - "repeat": optional, default 1, must be positive integer
//
// [ ] Open (create or truncate) file with given name
//
// [ ] Write content to file "repeat" times
//
// [ ] Close the file safely, handle write errors
//
// [ ] Write HTTP response:
//     - Status: 200 OK if success
//     - Body: confirmation message
//
// [ ] Handle possible errors:
//     - Missing or invalid parameters → 400 Bad Request
//     - File system errors → 500 Internal Server Error
//
// [ ] Sanitize filename to prevent directory traversal attacks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CreateFile crea o trunca el archivo indicado y escribe el contenido dado
// la cantidad de veces especificada. Retorna un mensaje de confirmación o error.
func CreateFile(name, content string, repeat int) (string, error) {
	// 1) Validar y sanitizar filename
	if name == "" {
		return "", fmt.Errorf("el parámetro 'name' es obligatorio")
	}
	// Prevenir directory traversal: solo permitimos nombres sin separadores
	if strings.ContainsAny(name, `/\`) || filepath.Base(name) != name {
		return "", fmt.Errorf("nombre de archivo inválido: %q", name)
	}

	// 2) Ajustar repeat si es inválido
	if repeat < 1 {
		repeat = 1
	}

	// 3) Crear o truncar el archivo
	f, err := os.Create(name)
	if err != nil {
		return "", fmt.Errorf("error al crear o truncar el archivo %q: %w", name, err)
	}
	defer f.Close()

	// 4) Escribir el contenido 'repeat' veces
	for i := 0; i < repeat; i++ {
		if _, err := f.WriteString(content); err != nil {
			return "", fmt.Errorf("error al escribir en el archivo %q: %w", name, err)
		}
	}

	// 5) Confirmación
	msg := fmt.Sprintf("Archivo %q creado/truncado con éxito (%d repeticiones)", name, repeat)
	return msg, nil
}
