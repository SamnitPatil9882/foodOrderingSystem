package app

import (
	"database/sql"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/category"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	repository "github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/dbquery"
)

type Dependencies struct {
	CategoryService category.Service
	FoodService     food.Service
}

func NewServices(db *sql.DB) Dependencies {

	//intialize repo dependencies
	categoryRepo := repository.NewCategoryRepo(db)
	foodRepo := repository.NewFoodRepo(db)
	//intialize service dependencies
	categoryService := category.NewService(categoryRepo)
	foodService := food.NewService(foodRepo)
	return Dependencies{
		CategoryService: categoryService,
		FoodService:     foodService,
	}
}
