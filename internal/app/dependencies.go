package app

import (
	"database/sql"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/category"
	repository "github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/dbquery"
)

type Dependencies struct {
	CategoryService category.Service
}

func NewServices(db *sql.DB) Dependencies {

	//intialize repo dependencies
	categoryRepo := repository.NewCategoryRepo(db)

	//intialize service dependencies
	categoryService := category.NewService(categoryRepo)
	return Dependencies{
		CategoryService: categoryService,
	}
}
