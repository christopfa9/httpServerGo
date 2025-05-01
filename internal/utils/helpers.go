// TODO: Implement helpers.go (Shared utility functions)
//
// [ ] Import necessary packages:
//     - fmt, strings, net, encoding/json, crypto/sha256, encoding/hex, etc.
//
// [ ] Define helper functions such as:
//
//     [ ] ParseQueryParams(rawQuery string) map[string]string
//         - Parses URL query string into key-value map
//
//     [ ] WriteHTTPResponse(conn net.Conn, statusCode int, contentType string, body string)
//         - Builds and writes a basic HTTP/1.0 response
//
//     [ ] SanitizeFileName(name string) string
//         - Prevents path traversal (removes "../", etc.)
//
//     [ ] JSONResponse(data any) (string, error)
//         - Converts a struct or map to a pretty-printed JSON string
//
//     [ ] SHA256Hash(input string) string
//         - Computes SHA-256 hash and returns it as hex string
//
// [ ] Ensure proper error handling and logging where needed
