package category

import "context"

type CategoryStorer interface {
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, categoryID int64) (Category, error)
	CreateCategory(ctx context.Context, category Category) error
	UpdateCategory(ctx context.Context, categoryID int64, category Category) error
}
type Category struct {
	ID           int
	CategoryName string
	description  string
	IsAcive      bool
}
