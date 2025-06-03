package commands

import (
	"strings"
)

// ToUpper convierte la cadena de texto recibida a mayúsculas y la devuelve.
// No retorna error ya que cualquier texto (incluso vacío) es válido.
func ToUpper(text string) (string, error) {
	return strings.ToUpper(text), nil
}
