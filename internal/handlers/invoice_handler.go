package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Juliannars1/invoice-service/internal/models"
	"github.com/Juliannars1/invoice-service/internal/service"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	service *service.InvoiceService
}

func NewInvoiceHandler(invoiceService *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		service: invoiceService,
	}
}
func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var invoice models.Invoice
	invoice.ID = "" 
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Now()
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateInvoice(r.Context(), &invoice)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Number := params["number"]
    if Number == "" {
		http.Error(w, "Missing invoice Number", http.StatusBadRequest)
		return
	}
	
	invoice, err := h.service.GetInvoice(r.Context(), Number)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	
	json.NewEncoder(w).Encode(invoice)
}
func (h *InvoiceHandler) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.service.GetAllInvoices(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get invoices", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(invoices)
}
func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	Number := params["number"]
	if Number == "" {
		http.Error(w, "Missing invoice NumberID", http.StatusBadRequest)
		return
	}

	
	var invoice map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateInvoice(r.Context(), Number, &invoice)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update invoice", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Number := params["number"]
	if Number == "" {
		http.Error(w, "Missing invoice NumberID", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteInvoice(r.Context(), Number)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete invoice", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
