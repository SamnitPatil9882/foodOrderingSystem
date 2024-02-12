package category

import (
	"context"
	"errors"
	"log"

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
	// UpdateCategoryStatus(ctx context.Context, categoryID int, updatedStatus int) (dto.Category, error)
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
	if categroyID <= 0 {
		return category, errors.New("invalid category id")
	}
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
	valres := validate(&createCategory)
	if !valres {
		return category, errors.New("invalid request details")
	}
	categoryDB, err := cs.categoryRepo.CreateCategory(ctx, createCategory)
	if err != nil {
		log.Println("error occred in create category service: " + err.Error())
	}

	category = MapRepoObjectToDto(categoryDB)
	return category, nil
}

func (cs *service) UpdateCategory(ctx context.Context, updateCategory dto.Category) (dto.Category, error) {

	// category := dto.Category{}

	categoryDB, err := cs.categoryRepo.UpdateCategory(ctx, updateCategory)
	if err != nil {
		log.Println("error occred in updating category service: " + err.Error())
		return dto.Category{},err
	}

	return categoryDB, nil
}

// func (cs *service) UpdateCategoryStatus(ctx context.Context, categoryID int, updatedStatus int) (dto.Category, error) {
// 	categoryDB, err := cs.categoryRepo.UpdateCategoryStatus(ctx, int64(categoryID), updatedStatus)
// 	ctgry := dto.Category{}

// 	if err != nil {
// 		log.Println("error occured in updating status: ", err.Error())
// 		return ctgry, err
// 	}

// 	ctgry = MapRepoObjectToDto(categoryDB)
// 	return ctgry, nil
// }
