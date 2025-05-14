package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func startTestServer(t *testing.T) (string, func()) {
	t.Helper()

	// Listen on a random free port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}
	port := fmt.Sprintf("%d", listener.Addr().(*net.TCPAddr).Port)
	listener.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	// Set PORT env variable for main
	os.Setenv("PORT", port)

	// Run server in goroutine
	go func() {
		defer wg.Done()
		main()
	}()

	time.Sleep(500 * time.Millisecond) // wait for server startup

	stop := func() {
		// This relies on main listening for OS signals; for tests you might need to improve shutdown.
		// For now just wait a bit and rely on test exit.
		time.Sleep(500 * time.Millisecond)
		wg.Wait()
	}

	return port, stop
}

func sendRequest(t *testing.T, port, req string) string {
	t.Helper()
	conn, err := net.DialTimeout("tcp", "localhost:"+port, 3*time.Second)
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(req))
	if err != nil {
		t.Fatalf("Failed to write request: %v", err)
	}

	reader := bufio.NewReader(conn)
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	return statusLine
}

func TestIntegration_AllEndpoints(t *testing.T) {
	port, stop := startTestServer(t)
	defer stop()

	tests := []struct {
		name    string
		request string
		want    string // Expected substring in status line or response line
	}{
		{"Fibonacci valid", "GET /fibonacci?num=10 HTTP/1.0\r\n\r\n", "200 OK"},
		{"CreateFile valid", "GET /createfile?name=testfile.txt&content=hello&repeat=1 HTTP/1.0\r\n\r\n", "200 OK"},
		{"DeleteFile valid", "GET /deletefile?name=testfile.txt HTTP/1.0\r\n\r\n", "200 OK"},
		{"Reverse valid", "GET /reverse?text=abcd HTTP/1.0\r\n\r\n", "200 OK"},
		{"ToUpper valid", "GET /toupper?text=abc HTTP/1.0\r\n\r\n", "200 OK"},
		{"Random valid", "GET /random?count=3&min=1&max=10 HTTP/1.0\r\n\r\n", "200 OK"},
		{"Timestamp valid", "GET /timestamp HTTP/1.0\r\n\r\n", "200 OK"},
		{"Hash valid", "GET /hash?text=test HTTP/1.0\r\n\r\n", "200 OK"},
		{"Simulate valid", "GET /simulate?seconds=0&task=test HTTP/1.0\r\n\r\n", "200 OK"},
		{"Sleep valid", "GET /sleep?seconds=0 HTTP/1.0\r\n\r\n", "200 OK"},
		{"LoadTest valid", "GET /loadtest?tasks=1&sleep=0 HTTP/1.0\r\n\r\n", "200 OK"},
		{"Status valid", "GET /status HTTP/1.0\r\n\r\n", "200 OK"},
		{"Help valid", "GET /help HTTP/1.0\r\n\r\n", "200 OK"},
		{"Unknown route", "GET /unknown HTTP/1.0\r\n\r\n", "404 Not Found"},
		{"Method not allowed", "POST /fibonacci?num=10 HTTP/1.0\r\n\r\n", "405 Method Not Allowed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := sendRequest(t, port, tt.request)
			if !strings.Contains(status, tt.want) {
				t.Errorf("Request %q: expected status containing %q, got %q", tt.request, tt.want, status)
			}
		})
	}
}
