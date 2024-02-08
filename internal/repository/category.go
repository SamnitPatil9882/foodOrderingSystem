package repository

import "context"

// "github.com/SamnitPatil9882/foodOrderingSystem/app/category"
type CategoryStorer interface {
	GetCategories(ctx context.Context) ([]Category, error)
	// GetCategory(ctx context.Context, categoryID int64) (Category, error)
	// CreateCategory(ctx context.Context, category Category) error
	// UpdateCategory(ctx context.Context, categoryID int64, category Category) error
}

type Category struct {
	ID           int
	CategoryName string
	Description  string
	IsAcive      bool
}
