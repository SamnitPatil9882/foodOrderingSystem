package food

import (
	"context"
	"fmt"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type service struct {
	foodRepo repository.FoodStorer
}
type Service interface {
	GetFoodList(ctx context.Context) ([]dto.Food, error)
	GetFoodListByCategory(ctx context.Context, categoryID int) ([]dto.Food, error)
	CreateFoodItem(ctx context.Context, fd dto.FoodCreateRequest) (dto.Food, error)
	UpdateFoodItem(ctx context.Context, fd dto.Food) (dto.Food, error)
	GetFoodByID(ctx context.Context, foodID int) (dto.Food, error)
}

func NewService(foodRepo repository.FoodStorer) Service {
	return &service{
		foodRepo: foodRepo,
	}
}

func (fdSvc *service) GetFoodList(ctx context.Context) ([]dto.Food, error) {
	fdList := make([]dto.Food, 0)

	fdListDB, err := fdSvc.foodRepo.GetListOfOrder(ctx)
	if err != nil {
		fmt.Println("error occured in getting data from db in service : " + err.Error())
	}
	for _, fd := range fdListDB {
		fdList = append(fdList, MapRepoObjectToDto(fd))
	}
	return fdList, nil
}
func (fdSvc *service) GetFoodListByCategory(ctx context.Context, categoryID int) ([]dto.Food, error) {
	fdList := make([]dto.Food, 0)

	// fdListDB, err := fdSvc.foodRepo.GetFoodByCategory(ctx, categoryID)
	fdListDB, err := fdSvc.foodRepo.GetFoodByCategory(ctx, categoryID)
	// GetFoodByCategory(ctx context.Context,categoryID int) ([]Food, error)
	if err != nil {
		fmt.Println("error occured in getting data from db in service : " + err.Error())
	}
	for _, fd := range fdListDB {
		fdList = append(fdList, MapRepoObjectToDto(fd))
	}
	return fdList, nil
}

func (fdSvc *service) CreateFoodItem(ctx context.Context, fd dto.FoodCreateRequest) (dto.Food, error) {

	resFd := dto.Food{}
	fdDB, err := fdSvc.foodRepo.CreateFood(ctx, fd)
	if err != nil {
		fmt.Println("error occured in getting data from db in service : " + err.Error())
		return resFd, err
	}
	resFd = MapRepoObjectToDto(fdDB)

	return resFd, nil
}

func (fdSvc *service) UpdateFoodItem(ctx context.Context, fd dto.Food) (dto.Food, error) {

	resFd := dto.Food{}
	fdDB, err := fdSvc.foodRepo.UpdateFood(ctx, fd)
	if err != nil {
		fmt.Println("error occured in getting data from db in service : " + err.Error())
		return resFd, err
	}
	resFd = MapRepoObjectToDto(fdDB)

	return resFd, nil
}
func (fdSvc *service) GetFoodByID(ctx context.Context, foodID int) (dto.Food, error) {
	resFd := dto.Food{}
	fdDB, err := fdSvc.foodRepo.GetFoodByID(ctx, int64(foodID))
	if err != nil {
		fmt.Println("error occured in getting data from db in service : " + err.Error())
		return resFd, err
	}
	resFd = MapRepoObjectToDto(fdDB)

	return resFd, nil
}
