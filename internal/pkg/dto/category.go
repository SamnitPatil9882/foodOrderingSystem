package dto

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	IsAcive      bool   `json:"is_active"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}
