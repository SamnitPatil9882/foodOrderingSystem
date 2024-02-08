package category

import (
	"context"
	"fmt"

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
	GetCategory(ctx context.Context, categoryID int64) (dto.Category, error)
	CreateCategory(ctx context.Context, createCategory dto.CategoryCreateRequest) (dto.Category, error)
	UpdateCategory(ctx context.Context, updateCategory dto.Category) (dto.Category, error)
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

func (cs *service) GetCategory(ctx context.Context, categroyID int64) (dto.Category, error) {
	category := dto.Category{}

	categoryDB, err := cs.categoryRepo.GetCategory(ctx, categroyID)
	if err != nil {
		return category, err
	}
	//GetCategory
	category = MapRepoObjectToDto(categoryDB)

	return category, nil
}
func (cs *service) CreateCategory(ctx context.Context, createCategory dto.CategoryCreateRequest) (dto.Category, error) {

	category := dto.Category{}

	categoryDB, err := cs.categoryRepo.CreateCategory(ctx, createCategory)
	if err != nil {
		fmt.Println("error occred in create category service: " + err.Error())
	}

	category = MapRepoObjectToDto(categoryDB)
	return category, nil
}

func (cs *service) UpdateCategory(ctx context.Context, updateCategory dto.Category) (dto.Category, error) {

	// category := dto.Category{}

	categoryDB, err := cs.categoryRepo.UpdateCategory(ctx, updateCategory)
	if err != nil {
		fmt.Println("error occred in updating category service: " + err.Error())
	}

	return categoryDB, nil
}
