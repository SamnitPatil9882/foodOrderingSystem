package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type service struct {
	Cart        []dto.CartItem
	foodService food.Service
	invoiceRepo repository.InvoiceStorer
}
type Service interface {
	AddOrderItem(ctx context.Context, orderItemId int, quantity int) error
	GetOrderList(ctx context.Context) []dto.CartItem
	RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error)
	CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error)
}

func NewService(orderSvc food.Service, invoiceRepo repository.InvoiceStorer) Service {
	return &service{
		foodService: orderSvc,
		invoiceRepo: invoiceRepo,
	}
}
func (ods *service) AddOrderItem(ctx context.Context, orderItemId int, quantity int) error {

	flg := isIDPresent(ods.Cart, orderItemId, quantity)
	if flg {
		// ods.Cart = append(ods.Cart, dto.CartItem{ID: orderItemId, Quantity: quantity, FoodName: fd.Name, Price: quantity * int(fd.Price)})
		return nil
	}
	fd, err := ods.foodService.GetFoodByID(ctx, orderItemId)
	if err != nil {
		return errors.New("food id not matched: " + err.Error())
	}

	ods.Cart = append(ods.Cart, dto.CartItem{ID: orderItemId, Quantity: quantity, FoodName: fd.Name, Price: quantity * int(fd.Price)})

	fmt.Println(ods.Cart)
	return nil
}
func (ods *service) GetOrderList(ctx context.Context) []dto.CartItem {

	return ods.Cart
}

func (ods *service) RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error) {
	flg := isIDPresent(ods.Cart, id, 1)
	if !flg {
		return ods.Cart, errors.New("bad request")
	}
	var err error
	ods.Cart, err = removeItemFromCart(ods.Cart, id)
	return ods.Cart, err
}
func (ods *service) CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error) {
	invoiceInfo.CartItem = ods.Cart
	res, err := ods.invoiceRepo.CreateInvoice(ctx, invoiceInfo)

	if err != nil {
		return dto.Invoice{}, err
	}

	return res, nil
}
func isIDPresent(cart []dto.CartItem, id int, quantity int) bool {
	for idx := range cart {
		if cart[idx].ID == id {
			if cart[idx].Quantity == 0 {
				cart[idx].Quantity = 1
			}
			cart[idx].Price = quantity * (cart[idx].Price / cart[idx].Quantity)
			cart[idx].Quantity = quantity

			return true // ID found in cart
		}
	}
	return false // ID not found in cart
}
func removeItemFromCart(cart []dto.CartItem, idToRemove int) ([]dto.CartItem, error) {
	for i, item := range cart {
		if item.ID == idToRemove {
			cart = append(cart[:i], cart[i+1:]...)
			return cart, nil
		}
	}

	return cart, errors.New("id not found")
}
