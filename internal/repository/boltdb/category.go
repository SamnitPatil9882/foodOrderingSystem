package boltdb

import (
	"context"
	"database/sql"
	"fmt"

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
