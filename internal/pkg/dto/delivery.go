package dto

type Delivery struct {
	ID        int    `json:"id"`
	OrderID   int    `json:"order_id"`
	UserID    int    `json:"user_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
}

type DeliveryUpdateRequst struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	// EndTime string `json:"end_time"`
	Status  string `json:"status"`
}
