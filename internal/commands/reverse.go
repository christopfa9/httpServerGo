package commands

// TODO: Implement reverse.go (Handles /reverse?text=)
//
// [ ] Import necessary packages:
//     - fmt, net, strings
//
// [ ] Define function HandleReverse(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "text" parameter:
//     - Required
//     - Return 400 Bad Request if missing
//
// [ ] Reverse the input string:
//     - Use rune slice to handle UTF-8 properly
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: reversed string
//
// [ ] Handle missing parameter or internal errors
//
// [ ] Log request and response (optional)

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
