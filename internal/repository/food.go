package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

type FoodStorer interface {
	GetListOfOrder(ctx context.Context) ([]Food, error)
	GetFoodByCategory(ctx context.Context, categoryID int) ([]Food, error)
	// GetFoodByID(ctx context.Context, FoodID int64) (Food, error)
	UpdateFood(ctx context.Context, food dto.Food) (Food, error)
	CreateFood(ctx context.Context, food dto.FoodCreateRequest) (Food, error)
}

type Food struct {
	ID         int64
	CategoryID int64
	Price      int64
	Name       string
	IsVeg      int
}
