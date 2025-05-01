package commands

// TODO: Implement help.go (Handles /help)
//
// [ ] Import necessary packages:
//     - fmt, net, strings
//
// [ ] Define function HandleHelp(conn net.Conn)
//
// [ ] Build a help message listing all available endpoints:
//     - /fibonacci?num=N
//     - /createfile?name=...&content=...&repeat=...
//     - /deletefile?name=...
//     - /reverse?text=...
//     - /toupper?text=...
//     - /random?count=...&min=...&max=...
//     - /timestamp
//     - /hash?text=...
//     - /simulate?seconds=...&task=...
//     - /sleep?seconds=...
//     - /loadtest?tasks=...&sleep=...
//     - /status
//     - /help
//
// [ ] Write HTTP response:
//     - Status line: HTTP/1.0 200 OK
//     - Headers: Content-Type: text/plain
//     - Body: formatted list of commands with brief descriptions
//
// [ ] Ensure formatting is readable and output is concise
