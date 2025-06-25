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
	// 0) Initialize all worker pools (includes computepi, pow, etc.)
	server.InitWorkerPools()

	// 1) Read port from PORT environment variable or use 8080 by default
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("Worker starting on port %s", port)

	// 2) Start the listener in a goroutine (blocks inside StartListener)
	go func() {
		if err := server.StartListener(port); err != nil {
			log.Fatalf("Worker error: %v", err)
		}
	}()

	// 3) Capture signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Worker received %s, shutting down...", sig)

	// (Optional) Give a few seconds for active connections to finish
	time.Sleep(1 * time.Second)
}
