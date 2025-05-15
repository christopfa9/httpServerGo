package commands

// Reverse invierte la cadena de texto recibida y la devuelve.
// No retorna error ya que cualquier texto (incluso vacío) es válido.
func Reverse(text string) (string, error) {
	// Convertimos a slice de runas para manejar UTF-8 correctamente
	runes := []rune(text)
	// Intercambiamos elementos desde los extremos hacia el centro
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}
