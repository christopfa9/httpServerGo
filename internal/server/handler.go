package server

// TODO: Implement handler.go (Handles and routes HTTP requests)
//
// [ ] Import necessary packages:
//     - fmt, io, net, strings, encoding/json, time
//     - internal/commands, internal/status
//
// [ ] Define a function HandleConnection(conn net.Conn)
//
// [ ] Read raw HTTP request data from conn
//
// [ ] Parse HTTP/1.0 request line and query parameters
//     - Only handle GET method
//     - Extract path and query (e.g., /fibonacci?num=5)
//
// [ ] Match path to supported commands:
//     - Call appropriate function from internal/commands
//
// [ ] Generate appropriate HTTP response:
//     - Status line (e.g., HTTP/1.0 200 OK)
//     - Headers (Content-Type, Content-Length, etc.)
//     - Body (text or JSON depending on route)
//
// [ ] Handle malformed requests:
//     - Return 400 Bad Request for invalid input
//     - Return 404 Not Found for unknown routes
//     - Return 500 Internal Server Error on panic
//
// [ ] Close connection after responding
//
// [ ] Update metrics (total connections, active handlers, etc.)
