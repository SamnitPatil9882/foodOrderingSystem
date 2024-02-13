package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

type InvoiceStorer interface {
	CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error)
	GetInvoiceByOrderID(ctx context.Context, orderId int, userId int, role string) (dto.Invoice, error)
}
type Invoice struct {
	ID            int
	OrderID       int
	PaymentMethod string
	CreatedAt     string
}
