package dto

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAcive     int    `json:"is_active"`
}

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type CategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAcive     int    `json:"is_active"`
}
