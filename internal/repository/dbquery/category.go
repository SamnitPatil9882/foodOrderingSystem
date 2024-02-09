package boltdb

import (
	"context"
	"database/sql"
	"fmt"
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
		fmt.Println("error occured in selecting categories: " + err.Error())
		return ctgryList, err
	}

	for rows.Next() {
		ctgry := repository.Category{}
		rows.Scan(&ctgry.ID, &ctgry.CategoryName, &ctgry.Description, &ctgry.IsAcive)
		fmt.Println(ctgry)
		ctgryList = append(ctgryList, ctgry)
	}
	return ctgryList, nil

}

func (cts *categoryStore) GetCategory(ctx context.Context, categoryID int64) (repository.Category, error) {

	category := repository.Category{
		ID:           0,
		CategoryName: "",
		Description:  "",
		IsAcive:      0,
	}
	query := "SELECT * FROM category WHERE id = " + strconv.FormatInt(categoryID, 10)

	rows, err := cts.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting category: " + err.Error())
		return category, err
	}
	for rows.Next() {
		// ctgry := repository.Category
		rows.Scan(&category.ID, &category.CategoryName, &category.Description, &category.IsAcive)
		fmt.Println(category)
		// ctgryList = append(ctgryList, ctgry)
	}
	fmt.Println("In db category", categoryID)
	// row.Scan(&category.ID, &category.CategoryName, &category.Description, &category.IsAcive)
	fmt.Println(category)
	return category, nil

}

func (cts *categoryStore) CreateCategory(ctx context.Context, category dto.CategoryCreateRequest) (repository.Category, error) {
	ctgry := repository.Category{}
	query := "INSERT INTO category (category_name,description,is_active) VALUES(?,?,?)"

	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("error occured in inseting values in db")
		return ctgry, err
	}
	res, err := statement.Exec(category.CategoryName, category.Description, category.IsAcive)
	if err != nil {
		return ctgry, err
	}
	categoryID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("erro occured in fetching inserted id")
		return ctgry, err
	}
	ctgry.ID = int(categoryID)
	ctgry.CategoryName = category.CategoryName
	ctgry.Description = category.Description
	ctgry.IsAcive = category.IsAcive
	return ctgry, nil

}

func (cts *categoryStore) UpdateCategory(ctx context.Context, catgory dto.Category) (dto.Category, error) {

	// ctgry := repository.Category{}
	query := fmt.Sprintf("UPDATE category SET category_name = '%s', description = '%s', is_active = %d WHERE id = %d", catgory.CategoryName, catgory.Description, catgory.IsAcive, catgory.ID)
	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("error occured in updating category db: " + err.Error())
		return catgory, err
	}
	statement.Exec()

	return catgory, nil

}

// func (cts *categoryStore) UpdateCategoryStatus(ctx context.Context, categoryID int64, UpdatedStatus int) (repository.Category, error) {
// 	if categoryID > 0 {
// 		categoryID = 1
// 	}
// 	category := repository.Category{}
// 	query := fmt.Sprintf("UPDATE category SET is_active = %d WHERE id = %d ",UpdatedStatus,categoryID)
// 	statement, err := cts.BaseRepsitory.DB.Prepare(query)
// 	if err != nil {
// 		fmt.Println("error occured in updating category db: " + err.Error())
// 		return category, err
// 	}
// 	statement.Exec()

// 	query = fmt.Sprintf("SELECT * FROM category WHERE id = %d",categoryID)
// 	// statement, err = cts.BaseRepsitory.DB.Prepare(query)
// 	// if err != nil {
// 	// 	fmt.Println("error occured in updating category db: " + err.Error())
// 	// 	return category, err
// 	// }
// 	// statement.Exec()

// 	rows, err := cts.BaseRepsitory.DB.Query(query)
// 	if err != nil {
// 		fmt.Println("error occured in selecting category: " + err.Error())
// 		return category, err
// 	}
// 	for rows.Next() {
// 		// ctgry := repository.Category
// 		rows.Scan(&category.ID, &category.CategoryName, &category.Description, &category.IsAcive)
// 		fmt.Println(category)
// 		// ctgryList = append(ctgryList, ctgry)
// 	}
// 	return category,nil
// }

/*
package repository

import (
	"fmt"

	// "github.com/SamnitPatil9882/foodOrderingSystem/app/category"
)

func GetCategory() ([]category.Category, error) {

	ctgryList := []category.Category{}
	query := "SELECT * FROM category"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting categories: " + err.Error())
		return nil, err
	}

	for rows.Next() {
		ctgry := category.Category{}
		rows.Scan(&ctgry.ID, &ctgry.CategoryName, &ctgry.Description, &ctgry.IsAcive)
		fmt.Println(ctgry)
		ctgryList = append(ctgryList, ctgry)
	}
	return ctgryList, nil

}

*/
