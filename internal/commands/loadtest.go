// TODO: Implement loadtest.go (Handles /loadtest?tasks=&sleep=)
//
// [ ] Import necessary packages:
//     - fmt, net, strconv, sync, time
//
// [ ] Define function HandleLoadTest(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "tasks": required, integer > 0
//     - "sleep": required, integer ≥ 0 (seconds)
//
// [ ] Use sync.WaitGroup to launch "tasks" number of goroutines
//     - Each goroutine should call time.Sleep(sleepSeconds)
//
// [ ] Wait for all goroutines to finish before responding
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: confirmation message (e.g., "Executed N concurrent tasks sleeping X seconds each")
//
// [ ] Handle malformed input with 400 Bad Request
//
// [ ] Log or track concurrent load (optional)
