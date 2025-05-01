package commands

// TODO: Implement deletefile.go (Handles /deletefile?name=)
//
// [ ] Import necessary packages:
//     - fmt, net, os, strings
//
// [ ] Define function HandleDeleteFile(conn net.Conn, params map[string]string)
//
// [ ] Extract and validate "name" parameter:
//     - Required
//     - Sanitize to prevent path traversal (no "../" etc.)
//
// [ ] Check if the file exists
//     - If not, respond with 404 Not Found
//
// [ ] Attempt to delete the file using os.Remove
//
// [ ] Write appropriate HTTP response:
//     - 200 OK if successfully deleted
//     - 500 Internal Server Error if deletion fails
//
// [ ] Handle missing or invalid parameters with 400 Bad Request
//
// [ ] Log result (optional)
