package main

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

	log.Printf("Server starting on port %s", port)

	// 2) Arrancamos el listener en una goroutine (bloquea dentro de StartListener)
	go func() {
		if err := server.StartListener(port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 3) Capturamos señales para shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Received %s, shutting down...", sig)

	// 4) Pequeña espera para que terminen handlers en vuelo (opcional)
	time.Sleep(500 * time.Millisecond)
	log.Println("Shutdown complete")
}
