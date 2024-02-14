package order

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	foodmock "github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food/mocks"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	repo "github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_service_AddOrderItem(t *testing.T) {
	// AddOrderItem(ctx context.Context, orderItemId int, quantity int) error
	tests := []struct {
		name        string
		orderItemId int
		quantity    int
		wantErr     bool
		setup       func(foodService *foodmock.Service)
	}{
		{
			name:    "successful",
			wantErr: false,
			setup: func(foodService *foodmock.Service) {
				foodService.On("GetFoodByID", context.Background(), mock.Anything).Return(dto.Food{
					ID:         1,
					CategoryID: 1,
					Price:      100,
					Name:       "Paneer",
					IsVeg:      1,
					IsAvail:    1,
				}, nil)
			},
		},
		{

			name:    "failure",
			wantErr: true,
			setup: func(foodService *foodmock.Service) {
				foodService.On("GetFoodByID", context.Background(), mock.Anything).Return(dto.Food{}, errors.New("error occured"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// fdRepo := &repo.FoodStorer{}
			// tt.setup(fdRepo)

			// fdRepo := &repo.FoodStorer{}
			invoiceRepo := &repo.InvoiceStorer{}
			deliveryRepo := &repo.DeliveryStorer{}
			orderRepo := &repo.OrderStorer{}
			fdSrv := foodmock.NewService(t)
			tt.setup(fdSrv)
			orderSrv := NewService(fdSrv, invoiceRepo, deliveryRepo, orderRepo)
			err := orderSrv.AddOrderItem(context.Background(), tt.orderItemId, tt.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_service_GetOrderList(t *testing.T) {
	// GetOrderList(ctx context.Context) []dto.CartItem
	tests := []struct {
		name string
		want []dto.CartItem
	}{
		{
			name: "successful",
			want: []dto.CartItem{
				{
					ID: 1,
					Quantity: 10,
					FoodName: "abc",
					Price: 1000,
				},
			},
			// setup func(ordSvc *service){

			// }
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// fdRepo := &repo.FoodStorer{}
			// tt.setup(fdRepo)

			// fdRepo := &repo.FoodStorer{}
			invoiceRepo := &repo.InvoiceStorer{}
			deliveryRepo := &repo.DeliveryStorer{}
			orderRepo := &repo.OrderStorer{}
			fdSrv := foodmock.NewService(t)
			orderSrv := NewService(fdSrv, invoiceRepo, deliveryRepo, orderRepo)
			fdSrv.On("GetFoodByID", context.Background(), mock.Anything).Return(dto.Food{
				ID:         1,
				CategoryID: 1,
				Price:      100,
				Name:       "abc",
				IsVeg:      1,
				IsAvail:    1,
			}, nil)

			orderSrv.AddOrderItem(context.Background(), 1, 10)
			got := orderSrv.GetOrderList(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFoodList() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_service_RemoveOrderItem(t *testing.T) {
	// RemoveOrderItem(ctx context.Context, id int) ([]dto.CartItem, error)
	tests := []struct {
		name    string
		ID      int
		want    []dto.CartItem
		wantErr bool
	}{
		{
			name: "successful",
			ID:   1,
			want: []dto.CartItem{
				dto.CartItem{
					ID: 1,
				},
			},
			// setup func(ordSvc *service){

			// }
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// fdRepo := &repo.FoodStorer{}
			// tt.setup(fdRepo)

			// fdRepo := &repo.FoodStorer{}
			invoiceRepo := &repo.InvoiceStorer{}
			deliveryRepo := &repo.DeliveryStorer{}
			orderRepo := &repo.OrderStorer{}
			fdSrv := foodmock.NewService(t)
			orderSrv := NewService(fdSrv, invoiceRepo, deliveryRepo, orderRepo)
			fdSrv.On("GetFoodByID", context.Background(), mock.Anything).Return(repository.Food{
				ID:         1,
				CategoryID: 1,
				Price:      100,
				Name:       "Paneer",
				IsVeg:      1,
				IsAvail:    1,
			}, nil)

			orderSrv.AddOrderItem(context.Background(), 1, 2)
			got := orderSrv.GetOrderList(context.Background())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFoodList() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateInvoice(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx         context.Context
		invoiceInfo dto.InvoiceCreation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.CreateInvoice(tt.args.ctx, tt.args.invoiceInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetDeliveryList(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.Delivery
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.GetDeliveryList(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetDeliveryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetDeliveryList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateDeliveryInfo(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx        context.Context
		updateInfo dto.DeliveryUpdateRequst
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			if err := ods.UpdateDeliveryInfo(tt.args.ctx, tt.args.updateInfo); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateDeliveryInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetListOfOrders(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx    context.Context
		userID int
		role   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.GetListOfOrders(tt.args.ctx, tt.args.userID, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetListOfOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetListOfOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetOrderByID(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx     context.Context
		orderID int
		userID  int
		role    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.GetOrderByID(tt.args.ctx, tt.args.orderID, tt.args.userID, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetOrderByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetOrderByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetInvoiceByOrderID(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx     context.Context
		orderID int
		userID  int
		role    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.GetInvoiceByOrderID(tt.args.ctx, tt.args.orderID, tt.args.userID, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetInvoiceByOrderID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetInvoiceByOrderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetOrderItemsByOrderID(t *testing.T) {
	type fields struct {
		Cart         []dto.CartItem
		foodService  food.Service
		invoiceRepo  repository.InvoiceStorer
		deliveryRepo repository.DeliveryStorer
		orderRepo    repository.OrderStorer
	}
	type args struct {
		ctx     context.Context
		orderID int
		role    string
		userID  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.CartItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ods := &service{
				Cart:         tt.fields.Cart,
				foodService:  tt.fields.foodService,
				invoiceRepo:  tt.fields.invoiceRepo,
				deliveryRepo: tt.fields.deliveryRepo,
				orderRepo:    tt.fields.orderRepo,
			}
			got, err := ods.GetOrderItemsByOrderID(tt.args.ctx, tt.args.orderID, tt.args.role, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetOrderItemsByOrderID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetOrderItemsByOrderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isIDPresent(t *testing.T) {
	type args struct {
		cart     []dto.CartItem
		id       int
		quantity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIDPresent(tt.args.cart, tt.args.id, tt.args.quantity); got != tt.want {
				t.Errorf("isIDPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeItemFromCart(t *testing.T) {
	type args struct {
		cart       []dto.CartItem
		idToRemove int
	}
	tests := []struct {
		name    string
		args    args
		want    []dto.CartItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeItemFromCart(tt.args.cart, tt.args.idToRemove)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeItemFromCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeItemFromCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
