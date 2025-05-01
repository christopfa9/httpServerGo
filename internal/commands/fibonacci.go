package commands

// TODO: Implement fibonacci.go (Handles /fibonacci?num=N)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, strings
//
// [ ] Define function HandleFibonacci(conn net.Conn, params map[string]string)
//
// [ ] Extract "num" parameter from params
//     - Validate that it exists and is a valid integer
//     - Return 400 Bad Request if invalid or missing
//
// [ ] Implement a recursive function fibonacci(n int) int
//     - Optionally use memoization if performance is a concern
//
// [ ] Compute the Fibonacci number for n
//
// [ ] Write the HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: result as string
//
// [ ] Handle errors gracefully and return appropriate HTTP error codes
//
// [ ] Log request and result (optional)
