// TODO: Implement commands_test.go (Unit tests for each command handler)
//
// [ ] Import necessary packages:
//     - testing, net, bytes, bufio, strings, io, os, encoding/json, time
//
// [ ] Create mockConn type implementing net.Conn for handler testing
//
// [ ] Write unit tests for each command:
//     [ ] TestFibonacciHandler_ValidInput
//         - Input: num=10 → Expect: "55"
//
//     [ ] TestCreateFileHandler_Valid
//         - Create temp file, check contents and cleanup
//
//     [ ] TestDeleteFileHandler_Valid
//         - Create then delete temp file, check deletion
//
//     [ ] TestReverseHandler_Valid
//         - Input: text=abc → Expect: "cba"
//
//     [ ] TestToUpperHandler_Valid
//         - Input: text=abc → Expect: "ABC"
//
//     [ ] TestRandomHandler_Valid
//         - Input: count=3&min=1&max=10 → Expect: 3 valid ints in range
//
//     [ ] TestTimestampHandler
//         - Expect: valid RFC3339 timestamp string
//
//     [ ] TestHashHandler_Valid
//         - Input: text=hello → Expect: known SHA-256 hash
//
//     [ ] TestSimulateHandler
//         - Input: seconds=1 → Expect delay + confirmation
//
//     [ ] TestSleepHandler
//         - Input: seconds=1 → Expect delay + confirmation
//
//     [ ] TestLoadTestHandler
//         - Input: tasks=5&sleep=1 → Expect completion of all goroutines
//
//     [ ] TestHelpHandler
//         - Expect list of supported endpoints
//
// [ ] Use table-driven tests where applicable
//
// [ ] Ensure 90%+ test coverage across commands
