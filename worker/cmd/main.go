package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"worker/internal/server"
)

func main() {
	// 0) Inicializamos todos los pools de workers (incluye computepi, pow, etc.)
	server.InitWorkerPools()

	// 1) Leemos el puerto desde la variable de entorno PORT o usamos 8080 por defecto
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("Worker starting on port %s", port)

	// 2) Arrancamos el listener en una goroutine (bloquea dentro de StartListener)
	go func() {
		if err := server.StartListener(port); err != nil {
			log.Fatalf("Worker error: %v", err)
		}
	}()

	// 3) Capturamos señales para shutdown limpio
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Worker received %s, shutting down...", sig)

	// (Opcional) Dar unos segundos para que terminen las conexiones activas
	time.Sleep(1 * time.Second)
}
