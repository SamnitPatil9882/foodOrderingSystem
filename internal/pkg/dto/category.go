package dto

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	IsAcive      int    `json:"is_active"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	IsAcive      int    `json:"is_active"`
}
