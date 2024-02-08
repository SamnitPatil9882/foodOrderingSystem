package category

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

// func (ctgry *Category)getCategory(){

// }

type service struct {
	categoryRepo repository.CategoryStorer
}
type Service interface {
	GetCategories(ctx context.Context) ([]dto.Category, error)
}

func NewService(categoryRepo repository.CategoryStorer) Service {
	return &service{
		categoryRepo: categoryRepo,
	}
}
func (cs *service) GetCategories(ctx context.Context) ([]dto.Category, error) {
	categories := make([]dto.Category, 0)

	categoriesDB, err := cs.categoryRepo.GetCategories(ctx)
	if err != nil {
		return categories, err
	}

	for _, categoryInfo := range categoriesDB {
		categories = append(categories, MapRepoObjectToDto(categoryInfo))
	}

	return categories, nil
}
