package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := routes.setupRoutes()
	
	server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go func() {
        log.Println("Starting server on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error starting server: %v", err)
        }
    }()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    log.Println("Shutting down server...")
    server.Close()
    log.Println("Server gracefully stopped")
}
