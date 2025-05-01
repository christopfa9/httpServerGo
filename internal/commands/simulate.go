// TODO: Implement simulate.go (Handles /simulate?seconds=&task=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, time
//
// [ ] Define function HandleSimulate(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "seconds": required, must be integer ≥ 0
//     - "task": optional, can be used for logging or reporting
//
// [ ] Simulate task execution by sleeping for the specified number of seconds
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: confirmation message (e.g., "Simulated task: taskname for X seconds")
//
// [ ] Return 400 Bad Request on invalid input
//
// [ ] Handle internal errors and respond with 500 Internal Server Error
//
// [ ] Optionally log the task and sleep duration
