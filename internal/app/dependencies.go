package app

import (
	"database/sql"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/category"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/order"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/user"
	repository "github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/dbquery"
)

type Dependencies struct {
	CategoryService category.Service
	FoodService     food.Service
	UserService     user.Service
	OrderService    order.Service
}

func NewServices(db *sql.DB) Dependencies {

	//intialize repo dependencies
	categoryRepo := repository.NewCategoryRepo(db)
	foodRepo := repository.NewFoodRepo(db)
	userRepo := repository.NewUserRepo(db)
	invoiceRepo := repository.NewInvoiceRepo(db)
	//intialize service dependencies
	categoryService := category.NewService(categoryRepo)
	foodService := food.NewService(foodRepo)
	userService := user.NewService(userRepo)
	orderService := order.NewService(foodService, invoiceRepo)
	return Dependencies{
		CategoryService: categoryService,
		FoodService:     foodService,
		UserService:     userService,
		OrderService:    orderService,
	}
}
