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
