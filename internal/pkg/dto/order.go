package dto

type OrderItem struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type CartItem struct {
	ID       int    `json:"id"`
	Quantity int    `json:"quantity"`
	FoodName string `json:"foodname"`
	Price    int    `json:"price"`
}

type Order struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	TotalAmout int    `json:"total_amount"`
	Location   string `json:"location"`
}
