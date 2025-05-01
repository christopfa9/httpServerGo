// TODO: Implement sleep.go (Handles /sleep?seconds=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, time
//
// [ ] Define function HandleSleep(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "seconds" parameter:
//     - Required, must be integer ≥ 0
//     - Return 400 Bad Request if missing or invalid
//
// [ ] Perform sleep using time.Sleep for the given duration
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: confirmation message (e.g., "Slept for X seconds")
//
// [ ] Handle input errors and respond appropriately
//
// [ ] Optionally log the sleep operation
