package food

// import (
// 	"context"
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
// 	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
// 	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/mocks"
// 	"github.com/stretchr/testify/mock"
// )

// func Test_service_GetFoodList(t *testing.T) {

// 	tests := []struct {
// 		name     string
// 		wantFood []dto.Food
// 		wantErr  bool
// 		setup    func(foodRepo *mocks.FoodStorer)
// 	}{
// 		{
// 			name: "successful",
// 			wantFood: []dto.Food{
// 				{
// 					ID:         1,
// 					CategoryID: 1,
// 					Price:      100,
// 					Name:       "Paneer",
// 					IsVeg:      1,
// 					IsAvail:    1,
// 				},
// 			},
// 			wantErr: false,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetListOfOrder", context.Background()).Return([]repository.Food{
// 					{
// 						ID:         1,
// 						CategoryID: 1,
// 						Price:      100,
// 						Name:       "Paneer",
// 						IsVeg:      1,
// 						IsAvail:    1,
// 					},
// 				}, nil)
// 			},
// 		},
// 		{
// 			name:     "failure",
// 			wantFood: []dto.Food{},
// 			wantErr:  true,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetListOfOrder", context.Background()).Return([]repository.Food{}, errors.New("errror occured"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			fdRepo := &mocks.FoodStorer{}
// 			tt.setup(fdRepo)
// 			fdSrv := NewService(fdRepo)
// 			got, err := fdSrv.GetFoodList(context.Background())
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.wantFood) {
// 				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.wantFood)
// 			}
// 		})
// 	}
// }

// func Test_service_GetFoodListByCategory(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		categoryID int
// 		wantFood   []dto.Food
// 		wantErr    bool
// 		setup      func(foodRepo *mocks.FoodStorer)
// 	}{
// 		{
// 			name:       "successful",
// 			categoryID: 1,
// 			wantFood: []dto.Food{
// 				{
// 					ID:         1,
// 					CategoryID: 1,
// 					Price:      100,
// 					Name:       "Paneer",
// 					IsVeg:      1,
// 					IsAvail:    1,
// 				},
// 			},
// 			wantErr: false,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetListOfOrder", context.Background()).Return([]repository.Food{
// 					{
// 						ID:         1,
// 						CategoryID: 1,
// 						Price:      100,
// 						Name:       "Paneer",
// 						IsVeg:      1,
// 						IsAvail:    1,
// 					},
// 				}, nil)
// 			},
// 		},
// 		{
// 			name:       "failure",
// 			categoryID: 10,
// 			wantFood:   []dto.Food{},
// 			wantErr:    true,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetListOfOrder", context.Background()).Return([]repository.Food{}, errors.New("errror occured"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			fdRepo := &mocks.FoodStorer{}
// 			tt.setup(fdRepo)
// 			fdSrv := NewService(fdRepo)
// 			got, err := fdSrv.GetFoodList(context.Background())
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.wantFood) {
// 				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.wantFood)
// 			}
// 		})
// 	}
// }

// func Test_service_CreateFoodItem(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		createFood dto.FoodCreateRequest
// 		wantFood   dto.Food
// 		wantErr    bool
// 		setup      func(foodRepo *mocks.FoodStorer)
// 	}{
// 		{
// 			name: "successful",
// 			createFood: dto.FoodCreateRequest{
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantFood: dto.Food{
// 				ID:         1,
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantErr: false,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("CreateFood", context.Background(), mock.Anything).Return(repository.Food{

// 					ID:         1,
// 					CategoryID: 1,
// 					Price:      100,
// 					Name:       "Paneer",
// 					IsVeg:      1,
// 					IsAvail:    1,
// 				}, nil)
// 			},
// 		},
// 		{

// 			name: "failure",
// 			createFood: dto.FoodCreateRequest{
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantFood: dto.Food{},
// 			wantErr:  true,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("CreateFood", context.Background(), mock.Anything).Return(repository.Food{}, errors.New("errror occured"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			fdRepo := &mocks.FoodStorer{}
// 			tt.setup(fdRepo)
// 			fdSrv := NewService(fdRepo)
// 			got, err := fdSrv.CreateFoodItem(context.Background(), tt.createFood)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.wantFood) {
// 				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.wantFood)
// 			}
// 		})
// 	}
// }

// func Test_service_UpdateFoodItem(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		updateFood dto.Food
// 		wantFood   dto.Food
// 		wantErr    bool
// 		setup      func(foodRepo *mocks.FoodStorer)
// 	}{
// 		{
// 			name: "successful",
// 			updateFood: dto.Food{
// 				ID:         1,
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantFood: dto.Food{
// 				ID:         1,
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantErr: false,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("UpdateFood", context.Background(), mock.Anything).Return(repository.Food{
// 					ID:         1,
// 					CategoryID: 1,
// 					Price:      100,
// 					Name:       "Paneer",
// 					IsVeg:      1,
// 					IsAvail:    1,
// 				}, nil)
// 			},
// 		},
// 		{

// 			name: "failure",
// 			updateFood: dto.Food{
// 				ID:         1,
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantFood: dto.Food{},
// 			wantErr:  true,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("UpdateFood", context.Background(), mock.Anything).Return(repository.Food{}, errors.New("errror occured"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			fdRepo := &mocks.FoodStorer{}
// 			tt.setup(fdRepo)
// 			fdSrv := NewService(fdRepo)
// 			got, err := fdSrv.UpdateFoodItem(context.Background(), tt.updateFood)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.wantFood) {
// 				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.wantFood)
// 			}
// 		})
// 	}
// }

// func Test_service_GetFoodByID(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		ID       int
// 		wantFood dto.Food
// 		wantErr  bool
// 		setup    func(foodRepo *mocks.FoodStorer)
// 	}{
// 		{
// 			name: "successful",
// 			wantFood: dto.Food{
// 				ID:         1,
// 				CategoryID: 1,
// 				Price:      100,
// 				Name:       "Paneer",
// 				IsVeg:      1,
// 				IsAvail:    1,
// 			},
// 			wantErr: false,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetFoodByID", context.Background(), mock.Anything).Return(repository.Food{
// 					ID:         1,
// 					CategoryID: 1,
// 					Price:      100,
// 					Name:       "Paneer",
// 					IsVeg:      1,
// 					IsAvail:    1,
// 				}, nil)
// 			},
// 		},
// 		{

// 			name: "failure",
// 			wantFood: dto.Food{},
// 			wantErr:  true,
// 			setup: func(foodRepo *mocks.FoodStorer) {
// 				foodRepo.On("GetFood", context.Background(), mock.Anything).Return(repository.Food{}, errors.New("errror occured"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			fdRepo := &mocks.FoodStorer{}
// 			tt.setup(fdRepo)
// 			fdSrv := NewService(fdRepo)
// 			got, err := fdSrv.GetFoodByID(context.Background(), tt.ID)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.wantFood) {
// 				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.wantFood)
// 			}
// 		})
// 	}
// }
