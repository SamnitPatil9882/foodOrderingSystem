package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type categoryStore struct {
	BaseRepsitory
}

func NewCategoryRepo(db *sql.DB) repository.CategoryStorer {
	return &categoryStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (cts *categoryStore) GetCategories(ctc context.Context) ([]repository.Category, error) {

	ctgryList := make([]repository.Category, 0)
	query := "SELECT * FROM category"

	rows, err := cts.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println("error occurred in selecting categories: " + err.Error())
		return ctgryList, internal.ErrFailedToFetchCategories
	}
	defer rows.Close()
	for rows.Next() {
		ctgry := repository.Category{}
		rows.Scan(&ctgry.ID, &ctgry.Name, &ctgry.Description, &ctgry.IsActive)
		log.Println(ctgry)
		ctgryList = append(ctgryList, ctgry)
	}
	return ctgryList, nil

}

func (cts *categoryStore) GetCategory(ctx context.Context, categoryID int64) (repository.Category, error) {

	category := repository.Category{
		ID:          0,
		Name:        "",
		Description: "",
		IsActive:    0,
	}
	query := "SELECT * FROM category WHERE id = " + strconv.FormatInt(categoryID, 10)

	rows, err := cts.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println("error occurred in selecting category: " + err.Error())
		return category, internal.ErrFailedToFetchCategory
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&category.ID, &category.Name, &category.Description, &category.IsActive)
		log.Println(category)
		return category, nil
	}
	log.Println("In db category", categoryID)
	log.Println(category)
	return category, internal.ErrCategoryNotFound

}

func (cts *categoryStore) CreateCategory(ctx context.Context, category dto.CategoryCreateRequest) (repository.Category, error) {
	ctgry := repository.Category{}
	query := "INSERT INTO category (name,description,is_active) VALUES(?,?,?)"

	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occurred in inserting values in db")
		return ctgry, internal.ErrFailedToInsertCategory
	}
	defer statement.Close()
	res, err := statement.Exec(category.Name, category.Description, category.IsActive)
	if err != nil {
		return ctgry, internal.ErrFailedToInsertCategory
	}
	categoryID, err := res.LastInsertId()
	if err != nil {
		log.Println("error occurred in fetching inserted id")
		return ctgry, internal.ErrFailedToInsertCategory
	}
	ctgry.ID = int(categoryID)
	ctgry.Name = category.Name
	ctgry.Description = category.Description
	ctgry.IsActive = category.IsActive
	return ctgry, nil

}

func (cts *categoryStore) UpdateCategory(ctx context.Context, catgory dto.Category) (dto.Category, error) {

	query := fmt.Sprintf("UPDATE category SET name = '%s', description = '%s', is_active = %d WHERE id = %d", catgory.Name, catgory.Description, catgory.IsActive, catgory.ID)
	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occurred in updating category db: " + err.Error())
		return catgory, internal.ErrFailedToUpdateCategory
	}
	defer statement.Close()
	res, err := statement.Exec()
	if err != nil {
		log.Println(err)
		return catgory, internal.ErrFailedToUpdateCategory
	}

	noOfRawAffected, err := res.RowsAffected()
	if err != nil {
		return catgory, internal.ErrFailedToUpdateCategory
	}
	if noOfRawAffected == 0 {
		return catgory, errors.New("id not matched")
	}
	return catgory, nil

}
