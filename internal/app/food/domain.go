package food

import (
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

func MapRepoObjectToDto(repoObj repository.Food) dto.Food {

	return dto.Food{
		ID:         repoObj.ID,
		CategoryID: repoObj.CategoryID,
		Price:      repoObj.Price,
		Name:       repoObj.Name,
		IsVeg:      repoObj.IsVeg,
	}
}
