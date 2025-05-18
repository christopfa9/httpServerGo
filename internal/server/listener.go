package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	listener    net.Listener
	connections sync.Map
	shutdown    chan struct{}
)

// StartListener abre un socket TCP en el puerto indicado y acepta conexiones entrantes.
func StartListener(port string) error {
	addr := fmt.Sprintf(":%s", port)
	var err error
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("no se pudo iniciar el listener en %s: %w", addr, err)
	}
	defer listener.Close()
	log.Printf("✔ Servidor escuchando en %s", addr)

	shutdown = make(chan struct{})

	for {
		select {
		case <-shutdown:
			log.Println("🔌 Cerrando listener...")
			return nil
		default:
			conn, err := listener.Accept()
			if err != nil {
				// Ignorar error si el listener ya se cerró
				if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
					return nil
				}
				log.Printf("⚠ error al aceptar conexión: %v", err)
				continue
			}
			log.Printf("→ Nueva conexión desde %s", conn.RemoteAddr())

			// Manejar la conexión en paralelo
			go func(c net.Conn) {
				defer func() {
					c.Close()
					connections.Delete(c.RemoteAddr())
					log.Printf("🛑 Conexión cerrada: %s", c.RemoteAddr())
				}()
				HandleConnection(c)
			}(conn)
		}
	}
}

// Shutdown detiene el listener y cierra todas las conexiones activas.
func Shutdown() {
	if listener != nil {
		close(shutdown)
		listener.Close()
	}

	connections.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			conn.Close()
		}
		return true
	})
}
