package server_test

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"

	"worker/internal/server"
)

// Inicializa los worker pools antes de correr tests
func init() {
	server.InitWorkerPools()
}

// helper para simular petición y leer respuesta completa sin bloqueo
func doRequest(t *testing.T, rawRequest string) string {
	t.Helper()

	client, serverConn := net.Pipe()
	defer client.Close()
	defer serverConn.Close()

	// Ejecutar el handler en goroutine para no bloquear
	go func() {
		server.HandleConnection(serverConn) // No retorna valor, por eso no se asigna
	}()

	// Escribir TODO el request (bloqueante)
	_, err := client.Write([]byte(rawRequest))
	if err != nil {
		t.Fatalf("failed to write request: %v", err)
	}

	// Leer respuesta completa línea a línea hasta EOF o timeout
	reader := bufio.NewReader(client)
	var response strings.Builder

	client.SetReadDeadline(time.Now().Add(5 * time.Second))
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				break // timeout indica fin de lectura
			}
			if err.Error() == "EOF" {
				break
			}
			t.Fatalf("failed reading response: %v", err)
		}
		response.WriteString(line)

		// Opcional: cortar si detecta fin de headers HTTP (línea vacía)
		if line == "\r\n" {
			break
		}
	}

	return response.String()
}

func TestHandleConnection_Fibonacci_Valid(t *testing.T) {
	req := "GET /fibonacci?num=10 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Fibonacci_MissingParam(t *testing.T) {
	req := "GET /fibonacci HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "400") {
		t.Errorf("Expected 400 Bad Request, got %q", resp)
	}
}

func TestHandleConnection_CreateFile_Valid(t *testing.T) {
	req := "GET /createfile?name=testfile.txt&content=hello&repeat=1 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_DeleteFile_Valid(t *testing.T) {
	// Para evitar fallos, crear archivo antes de eliminar
	reqCreate := "GET /createfile?name=testfile.txt&content=hello&repeat=1 HTTP/1.0\r\n\r\n"
	doRequest(t, reqCreate)

	reqDelete := "GET /deletefile?name=testfile.txt HTTP/1.0\r\n\r\n"
	resp := doRequest(t, reqDelete)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Reverse_Valid(t *testing.T) {
	req := "GET /reverse?text=abcd HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Toupper_Valid(t *testing.T) {
	req := "GET /toupper?text=abc HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Random_Valid(t *testing.T) {
	req := "GET /random?count=3&min=1&max=10 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Timestamp_Valid(t *testing.T) {
	req := "GET /timestamp HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Hash_Valid(t *testing.T) {
	req := "GET /hash?text=test HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Simulate_Valid(t *testing.T) {
	req := "GET /simulate?seconds=0&task=test HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Sleep_Valid(t *testing.T) {
	req := "GET /sleep?seconds=0 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_LoadTest_Valid(t *testing.T) {
	req := "GET /loadtest?tasks=1&sleep=0 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Status_Valid(t *testing.T) {
	req := "GET /status HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_Help_Valid(t *testing.T) {
	req := "GET /help HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "200 OK") {
		t.Errorf("Expected 200 OK, got %q", resp)
	}
}

func TestHandleConnection_UnknownRoute(t *testing.T) {
	req := "GET /unknown HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "404 Not Found") {
		t.Errorf("Expected 404 Not Found, got %q", resp)
	}
}

func TestHandleConnection_MethodNotAllowed(t *testing.T) {
	req := "POST /fibonacci?num=10 HTTP/1.0\r\n\r\n"
	resp := doRequest(t, req)

	if !strings.Contains(resp, "405 Method Not Allowed") {
		t.Errorf("Expected 405 Method Not Allowed, got %q", resp)
	}
}
