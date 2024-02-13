package order

import (
	"context"
	"errors"
	"log"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type service struct {
	Cart         []dto.CartItem
	foodService  food.Service
	invoiceRepo  repository.InvoiceStorer
	deliveryRepo repository.DeliveryStorer
	orderRepo    repository.OrderStorer
}
type Service interface {
	AddOrderItem(ctx context.Context, orderItemId int, quantity int) error
	GetOrderList(ctx context.Context) []dto.CartItem

	RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error)
	CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error)
	GetDeliveryList(ctx context.Context, userID int) ([]dto.Delivery, error)
	UpdateDeliveryInfo(ctx context.Context, updateInfo dto.DeliveryUpdateRequst) error
	GetListOfOrders(ctx context.Context, userID int, role string) ([]dto.Order, error)

	GetOrderByID(ctx context.Context, orderID int, userID int, role string) (dto.Order, error)
	GetInvoiceByOrderID(ctx context.Context, orderID int, userID int, role string) (dto.Invoice, error)
	GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error)
}

func NewService(orderSvc food.Service, invoiceRepo repository.InvoiceStorer, deliveryRepo repository.DeliveryStorer, orderRepo repository.OrderStorer) Service {
	return &service{
		foodService:  orderSvc,
		invoiceRepo:  invoiceRepo,
		deliveryRepo: deliveryRepo,
		orderRepo:    orderRepo,
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

	log.Println(ods.Cart)
	return nil
}
func (ods *service) GetOrderList(ctx context.Context) []dto.CartItem {

	return ods.Cart
}

func (ods *service) RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error) {
	if len(ods.Cart) == 0 {
		return []dto.CartItem{}, errors.New("cart is empty")
	}
	flg := isIDPresent(ods.Cart, id, 1)
	if !flg {
		return ods.Cart, errors.New("bad request")
	}
	var err error
	ods.Cart, err = removeItemFromCart(ods.Cart, id)
	return ods.Cart, err
}
func (ods *service) CreateInvoice(ctx context.Context, invoiceInfo dto.InvoiceCreation) (dto.Invoice, error) {
	if len(ods.Cart) == 0 {
		return dto.Invoice{}, errors.New("cart is empty")
	}
	invoiceInfo.CartItem = ods.Cart

	log.Printf("cart items in service: %v", ods.Cart)
	res, err := ods.invoiceRepo.CreateInvoice(ctx, invoiceInfo)
	if err != nil {
		return dto.Invoice{}, err
	}
	ods.Cart = ods.Cart[:0]
	return res, nil
}
func (ods *service) GetDeliveryList(ctx context.Context, userID int) ([]dto.Delivery, error) {

	dlvryList, err := ods.deliveryRepo.GetDeliveryList(ctx, userID)
	if err != nil {
		log.Println(err)
		return []dto.Delivery{}, err
	}
	return dlvryList, nil
}

func (ods *service) UpdateDeliveryInfo(ctx context.Context, updateInfo dto.DeliveryUpdateRequst) error {

	return ods.deliveryRepo.UpdateDeliveryInfo(ctx, updateInfo)
}

func (ods *service) GetListOfOrders(ctx context.Context, userID int, role string) ([]dto.Order, error) {
	return ods.orderRepo.GetListOfOrder(ctx, userID, role)
}
func (ods *service) GetOrderByID(ctx context.Context, orderID int, userID int, role string) (dto.Order, error) {
	return ods.orderRepo.GetOrderByID(ctx, orderID, userID, role)
}
func (ods *service) GetInvoiceByOrderID(ctx context.Context, orderID int, userID int, role string) (dto.Invoice, error) {
	return ods.invoiceRepo.GetInvoiceByOrderID(ctx, orderID, userID, role)
}
func (ods *service) GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error) {
	return ods.orderRepo.GetOrderItemsByOrderID(ctx, orderID, role, userID)
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
