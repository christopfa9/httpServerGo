// TODO: Implement createfile.go (Handles /createfile?name=&content=&repeat=)
//
// [ ] Import necessary packages:
//     - fmt, net, os, strconv, strings
//
// [ ] Define function HandleCreateFile(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate parameters:
//     - "name": required, no path traversal (sanitize input)
//     - "content": required
//     - "repeat": optional, default 1, must be positive integer
//
// [ ] Open (create or truncate) file with given name
//
// [ ] Write content to file "repeat" times
//
// [ ] Close the file safely, handle write errors
//
// [ ] Write HTTP response:
//     - Status: 200 OK if success
//     - Body: confirmation message
//
// [ ] Handle possible errors:
//     - Missing or invalid parameters → 400 Bad Request
//     - File system errors → 500 Internal Server Error
//
// [ ] Sanitize filename to prevent directory traversal attacks
