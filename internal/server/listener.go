// TODO: Implement listener.go (Handles TCP listening and connection acceptance)
//
// [ ] Import necessary packages:
//     - fmt, log, net, sync
//
// [ ] Define a function StartListener(port string) error
//
// [ ] Use net.Listen("tcp", port) to open a TCP socket
//
// [ ] Use a loop to accept incoming connections:
//     - For each accepted conn, launch a goroutine
//     - Pass conn to a handler function (e.g., HandleConnection)
//
// [ ] Handle and log errors when accepting connections
//
// [ ] Ensure safe shutdown with context or sync primitives (optional)
//
// [ ] Optionally track and update total connection count (sync-safe)
//
// [ ] Return any critical errors to the caller (main.go)
