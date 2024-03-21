package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
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
	log.Println("invoice info")
	log.Println(invoiceInfo)
	invoice := dto.Invoice{}
	currentTime := time.Now()
	createdAt := currentTime.Format("2006-01-02 15:04:05")
	log.Println("currenttime" + currentTime.String())
	var totalAmount int64
	for _, item := range invoiceInfo.CartItem {
		totalAmount += int64(item.Price)
	}

	//order
	query := "INSERT INTO `order` (user_id,created_at,total_amount,location) VALUES (?, ?, ?,?)"
	statement, err := invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement for inserting cart item: %v", err)
		return dto.Invoice{}, err
	}
	defer statement.Close()
	res, err := statement.Exec(invoiceInfo.UserID, createdAt, totalAmount, invoiceInfo.Location)
	if err != nil {
		log.Printf("Error inserting order: %v", err)
		return invoice, fmt.Errorf("failed to insert order: %w", err)
	}

	orderId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting orderId: %v", err)
		return invoice, fmt.Errorf("failed to get order ID: %w", err)
	}

	query = "INSERT INTO orderitem (id,order_id,food_id,quantity) VALUES (?,?,?,?)"
	statement, err = invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement for inserting cart item: %v", err)
	}
	defer statement.Close()
	for ind, item := range invoiceInfo.CartItem {
		log.Printf("orderid: %d", orderId)
		_, err := statement.Exec(ind+1, orderId, item.ID, item.Quantity)
		if err != nil {
			log.Printf("Error inserting cart item: %v", err)
			return invoice, fmt.Errorf("failed to insert cart item: %w", err)
		}
	}

	query = "INSERT INTO invoice (order_id,payment_method,created_at) VALUES(?,?,?)"
	statement, err = invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement for inserting invoice: %v", err)
		return dto.Invoice{}, err
	}
	res, err = statement.Exec(orderId, invoiceInfo.PaymentMethod, createdAt)
	if err != nil {
		log.Printf("Error inserting invoice: %v", err)
		return invoice, fmt.Errorf("failed to insert invoice: %w", err)
	}
	invoiceId, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error getting invoiceID: %v", err)
		return invoice, fmt.Errorf("failed to get invoice ID: %w", err)
	}
	invoice.ID = int(invoiceId)
	invoice.CreatedAt = createdAt
	invoice.OrderID = int(orderId)
	invoice.PaymentMethod = invoiceInfo.PaymentMethod

	query = "INSERT INTO delivery (order_id,start_at,status) VALUES(?,?,?)"
	statement, err = invst.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement for inserting delivery: %v", err)
		return dto.Invoice{}, err
	}
	_, err = statement.Exec(orderId, createdAt, "preparing")
	if err != nil {
		log.Printf("Error inserting delivery: %v", err)
		return invoice, fmt.Errorf("failed to insert delivery: %w", err)
	}
	return invoice, nil
}

func (invst *InvoiceStore) GetInvoiceByOrderID(ctx context.Context, orderId int, userId int, role string) (dto.Invoice, error) {
	inv := dto.Invoice{}
	if !invst.invoiceIDExists(orderId) {
		return inv, internal.ErrOrderIdNotExists
	}
	var query string
	if role == "customer" {
		query = fmt.Sprintf(`SELECT i.id, i.order_id, i.payment_method, i.created_at
		FROM invoice i
		JOIN "order" o ON i.order_id = o.id
		WHERE i.order_id = %d AND o.user_id = %d`, orderId, userId)
	} else if role == "deliveryboy" {
		query = fmt.Sprintf(`SELECT i.id, i.order_id, i.payment_method, i.created_at
		FROM invoice i
		JOIN "order" o ON i.order_id = o.id
		JOIN delivery d ON o.id = d.order_id
		WHERE i.order_id = %d AND d.deliveryboy_id = %d`, orderId, userId)
	} else {
		query = fmt.Sprintf("SELECT * FROM invoice WHERE order_id=%d", orderId)
	}
	rows, err := invst.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println(err)
		return inv, fmt.Errorf("failed to fetch invoice: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&inv.ID, &inv.OrderID, &inv.PaymentMethod, &inv.CreatedAt)
		return inv, nil
	}
	return inv, errors.New("order ID not found")
}
func (invst *InvoiceStore) invoiceIDExists(invoiceID int) bool {
	query := "SELECT COUNT(*) FROM invoice WHERE id = ?"
	var count int
	err := invst.BaseRepsitory.DB.QueryRow(query, invoiceID).Scan(&count)
	if err != nil {
		log.Println("Error checking if delivery ID exists:", err)
		return false
	}
	return count > 0
}