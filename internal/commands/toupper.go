package commands

// TODO: Implement toupper.go (Handles /toupper?text=)
//
// [ ] Import necessary packages:
//     - fmt, net, strings
//
// [ ] Define function HandleToUpper(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "text" parameter:
//     - Required
//     - Return 400 Bad Request if missing
//
// [ ] Convert the text to uppercase using strings.ToUpper
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: uppercase result
//
// [ ] Handle invalid input or internal errors
//
// [ ] Log request and result (optional)
