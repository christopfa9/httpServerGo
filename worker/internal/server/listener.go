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

// StartListener opens a TCP socket on the specified port and accepts incoming connections.
func StartListener(port string) error {
	addr := fmt.Sprintf(":%s", port)
	var err error
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to start listener on %s: %w", addr, err)
	}
	defer listener.Close()
	log.Printf("Server listening on %s", addr)

	shutdown = make(chan struct{})

	for {
		select {
		case <-shutdown:
			log.Println("Closing listener...")
			return nil
		default:
			conn, err := listener.Accept()
			if err != nil {
				// Ignore error if the listener has already been closed
				if opErr, ok := err.(*net.OpError); ok && opErr.Err.Error() == "use of closed network connection" {
					return nil
				}
				log.Printf("Error accepting connection: %v", err)
				continue
			}
			log.Printf("New connection from %s", conn.RemoteAddr())

			// Handle the connection concurrently
			go func(c net.Conn) {
				defer func() {
					c.Close()
					connections.Delete(c.RemoteAddr())
					log.Printf("Connection closed: %s", c.RemoteAddr())
				}()
				HandleConnection(c)
			}(conn)
		}
	}
}

// Shutdown stops the listener and closes all active connections.
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
