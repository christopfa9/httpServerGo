package server

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

import (
	"fmt"
	"log"
	"net"
)

// StartListener abre un socket TCP en el puerto indicado y acepta conexiones entrantes.
// Por cada conexión, lanza una goroutine que delega el trabajo en HandleConnection.
func StartListener(port string) error {
	// 1. Abrir el listener
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("no se pudo iniciar el listener en %s: %w", addr, err)
	}
	defer listener.Close()
	log.Printf("✔ Servidor escuchando en %s", addr)

	// 2. Bucle de aceptación de conexiones
	for {
		conn, err := listener.Accept()
		if err != nil {
			// 3. Registrar errores no críticos y continuar
			log.Printf("⚠ error al aceptar conexión: %v", err)
			continue
		}
		log.Printf("→ Nueva conexión desde %s", conn.RemoteAddr())

		// 4. Despachar a la función de manejo en paralelo
		go func(c net.Conn) {
			defer c.Close()
			HandleConnection(c)
		}(conn)
	}
}
