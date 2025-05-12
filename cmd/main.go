package main

// TODO: Implement main.go (HTTP Server Entry Point)
//
// [ ] Import necessary packages:
//     - fmt, log, net, os, os/signal, syscall, time
//     - internal/server, internal/status
//
// [ ] Define server configuration (e.g., port, max connections)
//
// [ ] Initialize metrics and status tracking (uptime, PID, connections)
//
// [ ] Start listening on the configured TCP port using net.Listen
//
// [ ] Use goroutines to handle incoming connections concurrently
//
// [ ] Dispatch requests to handler logic in internal/server/handler.go
//
// [ ] Handle graceful shutdown on SIGINT or SIGTERM:
//     - Close listener
//     - Clean up child goroutines or resources
//
// [ ] Log server startup and shutdown info
//
// [ ] Handle and log critical errors (e.g., port unavailable, panic recovery)
//
// [ ] Print instructions or banner on startup (optional)

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"httpServerGo/internal/server"
)

func main() {

	// 0) Arrancamos los pools de workers
	server.InitWorkerPools()

	// 1) Leemos puerto de la variable de entorno o usamos 8080 por defecto
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	log.Printf("⚡ Server starting on port %s", port)

	// 2) Arrancamos el listener en una goroutine (bloquea dentro de StartListener)
	go func() {
		if err := server.StartListener(port); err != nil {
			log.Fatalf("🔥 Server error: %v", err)
		}
	}()

	// 3) Capturamos señales para shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("✋ Received %s, shutting down...", sig)

	// 4) Pequeña espera para que terminen handlers en vuelo (opcional)
	time.Sleep(500 * time.Millisecond)
	log.Println("✅ Shutdown complete")
}
