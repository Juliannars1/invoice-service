package api

import (
	"log"

	"github.com/Juliannars1/invoice-service/internal/handlers"
	"github.com/Juliannars1/invoice-service/internal/repository"
	"github.com/Juliannars1/invoice-service/internal/service"
	"github.com/Juliannars1/invoice-service/internal/utils"
	"github.com/gorilla/mux"
)
func setupRoutes() *mux.Router{
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
	return router
}