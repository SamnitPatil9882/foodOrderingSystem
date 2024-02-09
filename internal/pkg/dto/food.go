package dto

type Food struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"category_id"`
	Price      int64  `json:"price"`
	Name       string `json:"name"`
	IsVeg      int    `json:"is_veg"`
}
type FoodCreateRequest struct {
	CategoryID int64  `json:"category_id"`
	Price      int64  `json:"price"`
	Name       string `json:"name"`
	IsVeg      int    `json:"is_veg"`
}
