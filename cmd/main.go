package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Juliannars1/invoice-service/internal/handlers"
	"github.com/Juliannars1/invoice-service/internal/repository"
	"github.com/Juliannars1/invoice-service/internal/service"
	"github.com/Juliannars1/invoice-service/internal/utils"

	"github.com/gorilla/mux"
)

func main() {
	mongoClient, err := utils.GetMongoDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(nil)

	invoiceRepo := repository.NewInvoiceRepository(mongoClient, "invoices", "invoice")
	invoiceService := service.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)
    
	router := mux.NewRouter()
	router.HandleFunc("/invoices", invoiceHandler.CreateInvoice).Methods("POST")
	router.HandleFunc("/invoices", invoiceHandler.GetAllInvoices).Methods("GET")
	router.HandleFunc("/invoices/{number}", invoiceHandler.GetInvoice).Methods("GET")
	router.HandleFunc("/invoices/{number}", invoiceHandler.UpdateInvoice).Methods("PUT")
	router.HandleFunc("/invoices/{number}", invoiceHandler.DeleteInvoice).Methods("DELETE")

	// Configurar el servidor HTTP
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    // Ejecutar el servidor en segundo plano
    go func() {
        log.Println("Starting server on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Error starting server: %v", err)
        }
    }()

    // Esperar una señal de interrupción para apagar el servidor
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    log.Println("Shutting down server...")
    server.Close()
    log.Println("Server gracefully stopped")
}
