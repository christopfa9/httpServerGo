// TODO: Implement random.go (Handles /random?count=&min=&max=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, math/rand, time, encoding/json
//
// [ ] Define function HandleRandom(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "count": required, integer > 0
//     - "min": required, integer
//     - "max": required, integer, must be >= min
//
// [ ] Seed the random number generator using time.Now().UnixNano()
//
// [ ] Generate a slice of count random integers in [min, max]
//
// [ ] Marshal the result as a JSON array
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: application/json
//     - Body: JSON array of random numbers
//
// [ ] Handle validation errors and respond with 400 Bad Request
//
// [ ] Handle internal errors with 500 Internal Server Error
