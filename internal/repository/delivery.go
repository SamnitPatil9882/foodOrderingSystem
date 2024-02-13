package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

type DeliveryStorer interface {
	GetDeliveryList(ctx context.Context,userID int) ([]dto.Delivery, error)
	UpdateDeliveryInfo(ctx context.Context, updateInfo dto.DeliveryUpdateRequst) error
}
type Delivery struct {
	ID        int
	OrderID   int
	UserID    int
	StartTime string
	EndTime   string
	Status    string
}
