package food

import (
	"context"
	"reflect"
	"testing"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

func TestNewService(t *testing.T) {
	type args struct {
		foodRepo repository.FoodStorer
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.foodRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetFoodList(t *testing.T) {
	type fields struct {
		foodRepo repository.FoodStorer
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fdSvc := &service{
				foodRepo: tt.fields.foodRepo,
			}
			got, err := fdSvc.GetFoodList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetFoodList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFoodList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetFoodListByCategory(t *testing.T) {
	type fields struct {
		foodRepo repository.FoodStorer
	}
	type args struct {
		ctx        context.Context
		categoryID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dto.Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fdSvc := &service{
				foodRepo: tt.fields.foodRepo,
			}
			got, err := fdSvc.GetFoodListByCategory(tt.args.ctx, tt.args.categoryID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetFoodListByCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFoodListByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateFoodItem(t *testing.T) {
	type fields struct {
		foodRepo repository.FoodStorer
	}
	type args struct {
		ctx context.Context
		fd  dto.FoodCreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fdSvc := &service{
				foodRepo: tt.fields.foodRepo,
			}
			got, err := fdSvc.CreateFoodItem(tt.args.ctx, tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateFoodItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateFoodItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateFoodItem(t *testing.T) {
	type fields struct {
		foodRepo repository.FoodStorer
	}
	type args struct {
		ctx context.Context
		fd  dto.Food
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fdSvc := &service{
				foodRepo: tt.fields.foodRepo,
			}
			got, err := fdSvc.UpdateFoodItem(tt.args.ctx, tt.args.fd)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateFoodItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.UpdateFoodItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetFoodByID(t *testing.T) {
	type fields struct {
		foodRepo repository.FoodStorer
	}
	type args struct {
		ctx    context.Context
		foodID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.Food
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fdSvc := &service{
				foodRepo: tt.fields.foodRepo,
			}
			got, err := fdSvc.GetFoodByID(tt.args.ctx, tt.args.foodID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetFoodByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetFoodByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
