package dto

type Invoice struct {
	ID            int    `json:"id"`
	OrderID       int    `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	CreatedAt     string `json:"created_at"`
}
type InvoiceCreation struct {
	UserID        int        `json:"user_id"`
	CartItem      []CartItem `json:"cart"`
	PaymentMethod string     `json:"payment_method"`
	Location      string     `json:"location"`
}
