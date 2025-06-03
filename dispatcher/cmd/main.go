package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dispatcher/internal" // ← aquí apuntamos a dispatcher/internal (sin "/dispatcher")
)

func main() {
	// 1) Puerto donde escuchará el Dispatcher (por defecto 9090,
	//    o bien lo lee de la variable de entorno “PORT”).
	port := "9090"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	log.Printf("Dispatcher starting on port %s\n", port)

	// 2) Leemos lista de Workers desde la variable de entorno “WORKERS”,
	//    separada por comas: “host1:8080,host2:8080,host3:8080”.
	workerEnv := os.Getenv("WORKERS")
	if workerEnv == "" {
		log.Fatalf("Necesito la variable de entorno WORKERS (host:port,host:port,...)")
	}
	workers := internal.ParseWorkersEnv(workerEnv)
	if len(workers) == 0 {
		log.Fatalf("WORKERS no contiene direcciones válidas")
	}

	// 3) Inicializamos el balanceador y healthchecks
	internal.InitDispatcher(workers)

	// 4) Abrimos listener TCP
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error al iniciar listener en puerto %s: %v", port, err)
	}

	// 5) Arrancamos ciclo para aceptar conexiones entrantes
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error Accept(): %v", err)
				continue
			}
			// Delegamos la conexión al dispatcher
			go internal.HandleConnection(conn)
		}
	}()

	// 6) Manejo de cierre con señales OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Dispatcher received %s, shutting down...\n", sig)

	// 7) Permitimos un breve lapso para drenar conexiones activas
	time.Sleep(1 * time.Second)
	if err := listener.Close(); err != nil {
		log.Printf("Error cerrando listener: %v", err)
	}
	log.Println("Dispatcher stopped.")
}
