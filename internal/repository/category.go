package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

// "github.com/SamnitPatil9882/foodOrderingSystem/app/category"
type CategoryStorer interface {
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategory(ctx context.Context, categoryID int64) (Category, error)
	CreateCategory(ctx context.Context, category dto.CategoryCreateRequest) (Category, error)
	UpdateCategory(ctx context.Context, category dto.Category) (dto.Category, error)
	// UpdateCategoryStatus(ctx context.Context,categoryID int64,UpdatedStatus int)(Category, error)
	DelCategory(ctx context.Context, categoryID int64) error
}

type Category struct {
	ID          int
	Name        string
	Description string
	IsActive     int
}
