package dbquery

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type InvoiceStore struct {
	BaseRepsitory
}

func NewInvoiceRepo(db *sql.DB) repository.InvoiceStorer {
	return &InvoiceStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (invst *InvoiceStore) CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error) {
	log.Println(invoiceInfo)
	invoice := dto.Invoice{}
	currentTime := time.Now()
	createdAt := currentTime.Format("2006-01-02 15:04:05")
	var totalAmount int64
	for _, item := range invoiceInfo.CartItem {
		totalAmount += int64(item.Price)
	}

	//order
	query := "INSERT INTO `order` (user_id,created_at,total_amount,location) VALUES (?, ?, ?,?)"
	statement, err := invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error in prepareing for inserting cart item: %v", err)
	}
	defer statement.Close()
	res, err := statement.Exec(invoiceInfo.UserID, createdAt, totalAmount, invoiceInfo.Location)
	if err != nil {
		log.Printf("Error inserting order: %v", err)
		return invoice, err
	}

	orderId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting orderId: %v", err)
		return invoice, err
	}

	query = "INSERT INTO orderitem (id,order_id,food_id,quantity) VALUES (?,?,?,?)"
	statement, err = invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error in prepareing for inserting cart item: %v", err)
	}

	for ind, item := range invoiceInfo.CartItem {
		log.Printf("orderid: %d", orderId)
		_, err := statement.Exec(ind, orderId, item.ID, item.Quantity)
		if err != nil {
			log.Printf("Error inserting sp cart item: %v", err)
			return invoice, err
		}
	}

	query = "INSERT INTO invoice (order_id,payment_method,created_at) VALUES(?,?,?)"
	statement, err = invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error in prepareing for inserting invoice: %v", err)
	}
	res, err = statement.Exec(orderId, invoiceInfo.PaymentMethod, createdAt)
	if err != nil {
		log.Printf("Error inserting invoice: %v", err)
		return invoice, err
	}
	invoiceId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting invoiceID: %v", err)
		return invoice, err
	}
	invoice.ID = int(invoiceId)
	invoice.CreatedAt = createdAt
	invoice.OrderID = int(orderId)
	invoice.PaymentMethod = invoiceInfo.PaymentMethod

	return invoice, nil
}
