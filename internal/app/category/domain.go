package category

import (
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

func MapRepoObjectToDto(repoObj repository.Category) dto.Category {
	return dto.Category{
		ID:           repoObj.ID,
		CategoryName: repoObj.CategoryName,
		Description:  repoObj.Description,
		IsAcive:      repoObj.IsAcive,
	}
}
