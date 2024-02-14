package category

import (
	"context"
	"errors"
	"testing"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCategories(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		wantCategories []dto.Category
		wantErr        error
		setup          func(categoryRepo *mocks.CategoryStorer)
	}{
		{
			name: "Success",
			wantCategories: []dto.Category{
				{ID: 1, Name: "Category 1", Description: "Description 1", IsActive: 1},
				{ID: 2, Name: "Category 2", Description: "Description 2", IsActive: 1},
			},
			wantErr: nil,
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("GetCategories", mock.Anything).Return([]repository.Category{
					{ID: 1, Name: "Category 1", Description: "Description 1", IsActive: 1},
					{ID: 2, Name: "Category 2", Description: "Description 2", IsActive: 1},
				}, nil).Once()
			},
		},
		{
			name:           "Error",
			wantCategories: []dto.Category{},
			wantErr:        errors.New("repository error"),
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("GetCategories", mock.Anything).Return([]repository.Category{}, errors.New("repository error")).Once()
			},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			categoryRepo := &mocks.CategoryStorer{}

			// Apply setup function to mock repository
			if tt.setup != nil {
				tt.setup(categoryRepo)
			}

			// Create service with mock repository
			service := NewService(categoryRepo)

			// Call the GetCategories method
			categories, err := service.GetCategories(context.Background())

			// Check for error and compare results
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, tt.wantCategories, categories) {
				t.Errorf("unexpected categories: got %v, want %v", categories, tt.wantCategories)
			}

		})
	}
}

/*
func (cs *service) CreateCategory(ctx context.Context, createCategory dto.CategoryCreateRequest) (dto.Category, error) {

	category := dto.Category{}
	valres := validate(&createCategory)
	if !valres {
		return category, errors.New("invalid request details")
	}
	categoryDB, err := cs.categoryRepo.CreateCategory(ctx, createCategory)
	if err != nil {
		log.Println("error occred in create category service: " + err.Error())
		return dto.Category{}, err
	}

	category = MapRepoObjectToDto(categoryDB)
	return category, nil
}
*/

func TestCreateCategory(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		createCategory dto.CategoryCreateRequest
		wantCategory   dto.Category
		wantErr        error
		setup          func(categoryRepo *mocks.CategoryStorer)
	}{
		{
			name: "Success",
			createCategory: dto.CategoryCreateRequest{
				Name:        "Test Category",
				Description: "Test Description",
				IsActive:    1,
			},
			wantCategory: dto.Category{
				ID:          1,
				Name:        "Test Category",
				Description: "Test Description",
				IsActive:    1,
			},
			wantErr: nil,
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("CreateCategory", context.Background(), mock.Anything).Return(repository.Category{
					ID:          1,
					Name:        "Test Category",
					Description: "Test Description",
					IsActive:    1,
				}, nil).Once()
			},
		},
		{
			name: "Category already available",
			createCategory: dto.CategoryCreateRequest{
				Name:        "Category 1",
				Description: "Description 1",
				IsActive:    1,
			},
			wantCategory: dto.Category{},
			wantErr:      errors.New("category already available"),
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("CreateCategory", mock.Anything, mock.Anything).Return(repository.Category{}, errors.New("category already available")).Once()
			},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			categoryRepo := &mocks.CategoryStorer{}

			// Apply setup function to mock repository
			if tt.setup != nil {
				tt.setup(categoryRepo)
			}

			// Create service with mock repository
			service := NewService(categoryRepo)

			// Call the GetCategory method
			category, err := service.CreateCategory(context.Background(), tt.createCategory)

			// Check for error and compare results
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, tt.wantCategory, category) {
				t.Errorf("unexpected category: got %v, want %v", category, tt.wantCategory)
			}

			// Assert that all expectations were met
			categoryRepo.AssertExpectations(t)
		})
	}
}
func TestUpdateCategory(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		updateCategory dto.Category
		wantCategory   dto.Category
		wantErr        error
		setup          func(categoryRepo *mocks.CategoryStorer)
	}{
		{
			name: "Success",
			updateCategory: dto.Category{
				ID:1,
				Name:        "Test Category",
				Description: "Test Description",
				IsActive:    1,
			},
			wantCategory: dto.Category{
				ID:          1,
				Name:        "Test Category",
				Description: "Test Description",
				IsActive:    1,
			},
			wantErr: nil,
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("UpdateCategory", context.Background(), mock.Anything).Return(dto.Category{
					ID:          1,
					Name:        "Test Category",
					Description: "Test Description",
					IsActive:    1,
				}, nil).Once()
			},
		},
		{
			name: "Category already available",
			updateCategory: dto.Category{
				ID: 10,
				Name:        "Category 1",
				Description: "Description 1",
				IsActive:    1,
			},
			wantCategory: dto.Category{},
			wantErr:      errors.New("category not available"),
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("UpdateCategory", mock.Anything, mock.Anything).Return(dto.Category{}, errors.New("category already available")).Once()
			},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			categoryRepo := &mocks.CategoryStorer{}

			// Apply setup function to mock repository
			if tt.setup != nil {
				tt.setup(categoryRepo)
			}

			// Create service with mock repository
			service := NewService(categoryRepo)

			// Call the GetCategory method
			category, err := service.UpdateCategory(context.Background(), tt.updateCategory)

			// Check for error and compare results
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, tt.wantCategory, category) {
				t.Errorf("unexpected category: got %v, want %v", category, tt.wantCategory)
			}

			// Assert that all expectations were met
			categoryRepo.AssertExpectations(t)
		})
	}
}

func TestGetCategory(t *testing.T) {
	// Define test cases
	tests := []struct {
		name         string
		categoryID   int64
		wantCategory dto.Category
		wantErr      error
		setup        func(categoryRepo *mocks.CategoryStorer)
	}{
		{
			name:       "Success",
			categoryID: 1,
			wantCategory: dto.Category{
				ID:          1,
				Name:        "Test Category",
				Description: "Test Description",
				IsActive:    1,
			},
			wantErr: nil,
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("GetCategory", mock.Anything, int64(1)).Return(repository.Category{
					ID:          1,
					Name:        "Test Category",
					Description: "Test Description",
					IsActive:    1,
				}, nil).Once()
			},
		},
		{
			name:         "Invalid Category ID",
			categoryID:   0,
			wantCategory: dto.Category{},
			wantErr:      errors.New("invalid category id"),
			setup: func(categoryRepo *mocks.CategoryStorer) {
				// No setup required for this test case
			},
		},
		{
			name:         "Repository Error",
			categoryID:   2,
			wantCategory: dto.Category{},
			wantErr:      errors.New("repository error"),
			setup: func(categoryRepo *mocks.CategoryStorer) {
				categoryRepo.On("GetCategory", mock.Anything, int64(2)).Return(repository.Category{}, errors.New("repository error")).Once()
			},
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			categoryRepo := &mocks.CategoryStorer{}

			// Apply setup function to mock repository
			if tt.setup != nil {
				tt.setup(categoryRepo)
			}

			// Create service with mock repository
			service := NewService(categoryRepo)

			// Call the GetCategory method
			category, err := service.GetCategory(context.Background(), tt.categoryID)

			// Check for error and compare results
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("unexpected error: got %v, want %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, tt.wantCategory, category) {
				t.Errorf("unexpected category: got %v, want %v", category, tt.wantCategory)
			}

			// Assert that all expectations were met
			categoryRepo.AssertExpectations(t)
		})
	}
}
