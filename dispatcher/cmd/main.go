package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dispatcher/internal" // ← points to dispatcher/internal (no "/dispatcher")
)

func main() {
	// 1) Port where the Dispatcher will listen (default 9090,
	//    or read from the "PORT" environment variable).
	port := "9090"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("Dispatcher starting on port %s\n", port)

	// 2) Read list of Workers from the "WORKERS" environment variable,
	//    separated by commas: "host1:8080,host2:8080,host3:8080".
	workerEnv := os.Getenv("WORKERS")
	if workerEnv == "" {
		log.Fatalf("The WORKERS environment variable is required (host:port,host:port,...)")
	}
	workers := internal.ParseWorkersEnv(workerEnv)
	if len(workers) == 0 {
		log.Fatalf("WORKERS does not contain valid addresses")
	}

	// 3) Initialize dispatcher and health checks
	internal.InitDispatcher(workers)

	// 4) Open TCP listener
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start listener on port %s: %v", port, err)
	}

	// 5) Start loop to accept incoming connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Accept() error: %v", err)
				continue
			}
			// Delegate the connection to the dispatcher
			go internal.HandleConnection(conn)
		}
	}()

	// 6) Handle shutdown on OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Dispatcher received %s, shutting down...\n", sig)

	// 7) Allow a short time to drain active connections
	time.Sleep(1 * time.Second)
	if err := listener.Close(); err != nil {
		log.Printf("Error closing listener: %v", err)
	}
	log.Println("Dispatcher stopped.")
}
