package category

import (
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func MapRepoObjectToDto(repoObj repository.Category) dto.Category {
	return dto.Category{
		ID:          repoObj.ID,
		Name:        repoObj.Name,
		Description: repoObj.Description,
		IsAcive:     repoObj.IsAcive,
	}
}

func validate(createCategory *dto.CategoryCreateRequest) bool {
	if len(createCategory.Name) < 2 {
		return false
	}
	createCategory.Name = cases.Title(language.Und, cases.NoLower).String(createCategory.Name)
	createCategory.Description = cases.Title(language.Und, cases.NoLower).String(createCategory.Description)
	if createCategory.IsAcive < 0 {
		createCategory.IsAcive = 0
	} else if createCategory.IsAcive > 1 {
		createCategory.IsAcive = 1
	}
	return true
}
