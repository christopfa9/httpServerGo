package server_test

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"

	"httpServerGo/internal/server"
)

func init() {
	server.InitWorkerPools()
}

// helper sends raw HTTP request over net.Pipe and returns first response line (status)
func doRequest(t *testing.T, rawRequest string) string {
	t.Helper()
	client, serverConn := net.Pipe()
	defer client.Close()
	defer serverConn.Close()

	go func() {
		serverConn.SetDeadline(time.Now().Add(5 * time.Second))
		_, err := serverConn.Write([]byte(rawRequest))
		if err != nil {
			t.Fatalf("failed to write request: %v", err)
		}
	}()

	doneCh := make(chan struct{})
	go func() {
		server.HandleConnection(serverConn)
		close(doneCh)
	}()

	reader := bufio.NewReader(client)
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		t.Fatalf("failed to read response: %v", err)
	}

	<-doneCh
	return statusLine
}

func TestHandleConnection_Fibonacci_Valid(t *testing.T) {
	req := "GET /fibonacci?num=10 HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", status)
	}
}

func TestHandleConnection_Fibonacci_MissingParam(t *testing.T) {
	req := "GET /fibonacci HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "400") {
		t.Errorf("Expected 400 Bad Request, got %q", status)
	}
}

func TestHandleConnection_CreateFile_Valid(t *testing.T) {
	req := "GET /createfile?name=testfile.txt&content=hello&repeat=2 HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", status)
	}
}

func TestHandleConnection_DeleteFile_Valid(t *testing.T) {
	// Optionally create the file here or assume it exists
	req := "GET /deletefile?name=testfile.txt HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", status)
	}
}

func TestHandleConnection_UnknownRoute(t *testing.T) {
	req := "GET /unknown HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "404 Not Found") {
		t.Errorf("Expected 404 Not Found, got %q", status)
	}
}

func TestHandleConnection_MethodNotAllowed(t *testing.T) {
	req := "POST /fibonacci?num=10 HTTP/1.0\r\n\r\n"
	status := doRequest(t, req)
	if !strings.Contains(status, "405 Method Not Allowed") {
		t.Errorf("Expected 405 Method Not Allowed, got %q", status)
	}
}
