package service

import (
	"context"

	"github.com/Juliannars1/invoice-service/internal/models"
	"github.com/Juliannars1/invoice-service/internal/repository"
)

type InvoiceService struct {
	
	repo *repository.InvoiceRepository
}

func NewInvoiceService(repo *repository.InvoiceRepository) *InvoiceService {
	return &InvoiceService{
		repo: repo,
	}
}

func (svc *InvoiceService) CreateInvoice(ctx context.Context, invoice *models.Invoice) error {
	return svc.repo.CreateInvoice(ctx, invoice)
}

func (svc *InvoiceService) GetInvoice(ctx context.Context, Number string) (*models.Invoice, error) {
	invoice, err := svc.repo.GetInvoice(ctx, Number)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
func (svc *InvoiceService) GetAllInvoices(ctx context.Context) ([]*models.Invoice, error) {
	
	invoice, err := svc.repo.GetAllInvoices(ctx)
	
	if err != nil {
		return nil, err
	}

	return invoice.([]*models.Invoice), nil
}
func (svc *InvoiceService) UpdateInvoice(ctx context.Context, number string, invoice *map[string]interface{}) error {
	return svc.repo.UpdateInvoice(ctx, number, invoice)
}

func (svc *InvoiceService) DeleteInvoice(ctx context.Context, number string) error {
	return svc.repo.DeleteInvoice(ctx, number)
}
