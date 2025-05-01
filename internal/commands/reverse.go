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
