package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

type OrderStorer interface {
	GetListOfOrder(ctx context.Context,userId int,role string) ([]dto.Order, error)
	GetOrderByID(ctx context.Context, orderID int,userId int , role string) (dto.Order, error)
	GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error)
}
type Order struct {
	ID         int
	UserID     int
	CreatedAt  string
	TotalAmout int
	Location   string
}
