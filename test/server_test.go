// TODO: Implement server_test.go (Tests for listener and handler logic)
//
// [ ] Import necessary packages:
//     - testing, net, time, bufio, strings
//
// [ ] Setup test environment:
//     - Start the server in a goroutine (on a test port)
//     - Use net.Dial to send mock requests
//
// [ ] Write test cases:
//     [ ] TestServerStartup
//         - Ensure server starts without errors
//
//     [ ] TestHandleFibonacciValid
//         - Send /fibonacci?num=10 and validate response "55"
//
//     [ ] TestHandleNotFound
//         - Send unknown path and expect "404 Not Found"
//
//     [ ] TestMalformedRequest
//         - Send badly formatted HTTP request and expect "400 Bad Request"
//
// [ ] Use bufio.Reader to read server responses and assert expected output
//
// [ ] Clean up connections and shut down server cleanly
