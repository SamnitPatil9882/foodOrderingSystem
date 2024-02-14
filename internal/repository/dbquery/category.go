package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

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
		log.Println("error occured in selecting categories: " + err.Error())
		return ctgryList, err
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
		log.Println("error occured in selecting category: " + err.Error())
		return category, err
	}
	defer rows.Close()
	for rows.Next() {
		// ctgry := repository.Category
		rows.Scan(&category.ID, &category.Name, &category.Description, &category.IsActive)
		log.Println(category)
		// ctgryList = append(ctgryList, ctgry)
		return category, nil
	}
	log.Println("In db category", categoryID)
	// row.Scan(&category.ID, &category.Name, &category.Description, &category.IsActive)
	log.Println(category)
	return category, errors.New("category not found")

}

func (cts *categoryStore) CreateCategory(ctx context.Context, category dto.CategoryCreateRequest) (repository.Category, error) {
	ctgry := repository.Category{}
	query := "INSERT INTO category (name,description,is_active) VALUES(?,?,?)"

	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occured in inseting values in db")
		return ctgry, err
	}
	defer statement.Close()
	res, err := statement.Exec(category.Name, category.Description, category.IsActive)
	if err != nil {
		return ctgry, err
	}
	categoryID, err := res.LastInsertId()
	if err != nil {
		log.Println("erro occured in fetching inserted id")
		return ctgry, err
	}
	ctgry.ID = int(categoryID)
	ctgry.Name = category.Name
	ctgry.Description = category.Description
	ctgry.IsActive = category.IsActive
	return ctgry, nil

}

func (cts *categoryStore) UpdateCategory(ctx context.Context, catgory dto.Category) (dto.Category, error) {

	// ctgry := repository.Category{}
	query := fmt.Sprintf("UPDATE category SET name = '%s', description = '%s', is_active = %d WHERE id = %d", catgory.Name, catgory.Description, catgory.IsActive, catgory.ID)
	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occured in updating category db: " + err.Error())
		return catgory, err
	}
	defer statement.Close()
	res, err := statement.Exec()
	if err != nil {
		log.Println(err)
		return catgory, err
	}

	noOfRawAffected, err := res.RowsAffected()
	if err != nil {
		return catgory, err
	}
	if noOfRawAffected == 0 {
		return catgory, errors.New("id not matched")
	}
	return catgory, nil

}

// func (cts *categoryStore) UpdateCategoryStatus(ctx context.Context, categoryID int64, UpdatedStatus int) (repository.Category, error) {
// 	if categoryID > 0 {
// 		categoryID = 1
// 	}
// 	category := repository.Category{}
// 	query := log.Sprintf("UPDATE category SET is_active = %d WHERE id = %d ",UpdatedStatus,categoryID)
// 	statement, err := cts.BaseRepsitory.DB.Prepare(query)
// 	if err != nil {
// 		log.Println("error occured in updating category db: " + err.Error())
// 		return category, err
// 	}
// 	statement.Exec()

// 	query = log.Sprintf("SELECT * FROM category WHERE id = %d",categoryID)
// 	// statement, err = cts.BaseRepsitory.DB.Prepare(query)
// 	// if err != nil {
// 	// 	log.Println("error occured in updating category db: " + err.Error())
// 	// 	return category, err
// 	// }
// 	// statement.Exec()

// 	rows, err := cts.BaseRepsitory.DB.Query(query)
// 	if err != nil {
// 		log.Println("error occured in selecting category: " + err.Error())
// 		return category, err
// 	}
// 	for rows.Next() {
// 		// ctgry := repository.Category
// 		rows.Scan(&category.ID, &category.Name, &category.Description, &category.IsActive)
// 		log.Println(category)
// 		// ctgryList = append(ctgryList, ctgry)
// 	}
// 	return category,nil
// }

/*
package repository

import (
	"log"

	// "github.com/SamnitPatil9882/foodOrderingSystem/app/category"
)

func GetCategory() ([]category.Category, error) {

	ctgryList := []category.Category{}
	query := "SELECT * FROM category"

	rows, err := db.Query(query)
	if err != nil {
		log.Println("error occured in selecting categories: " + err.Error())
		return nil, err
	}

	for rows.Next() {
		ctgry := category.Category{}
		rows.Scan(&ctgry.ID, &ctgry.Name, &ctgry.Description, &ctgry.IsActive)
		log.Println(ctgry)
		ctgryList = append(ctgryList, ctgry)
	}
	return ctgryList, nil

}

*/
