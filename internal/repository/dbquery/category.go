package dbquery

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	fmt.Println("GetCategories")
	ctgryList := make([]repository.Category, 0)
	query := "SELECT * FROM category"

	rows, err := cts.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println("error occurred in selecting categories: " + err.Error())
		fmt.Println("GetCategories3")
		return ctgryList, internal.ErrFailedToFetchCategories
	}
	defer rows.Close()
	fmt.Println("GetCategories2")
	for rows.Next() {
		fmt.Println("GetCategories1")
		ctgry := repository.Category{}
		err = rows.Scan(&ctgry.ID, &ctgry.Name, &ctgry.Description, &ctgry.IsActive)
		if err != nil {
			log.Println(err)
			fmt.Println("GetCategories4")
			return []repository.Category{}, internal.ErrFailedToFetchCategories
		}
		// return []repository.Category{}, err
		log.Println(ctgry)
		ctgryList = append(ctgryList, ctgry)
	}

	fmt.Println("categrylist: ", ctgryList)
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
	query := "INSERT INTO category (name, description, is_active) VALUES (?, ?, ?)"

	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occurred in preparing the insert statement:", err.Error())
		return ctgry, internal.ErrFailedToInsertCategory
	}
	defer statement.Close()

	// Execute the insert statement
	res, err := statement.Exec(category.Name, category.Description, category.IsActive)
	if err != nil {
		// Check if the error is due to unique constraint violation
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println("error: category name already exists")
			return ctgry, internal.ErrCategoryNameExists
		}
		log.Println("error occurred in executing the insert statement:", err.Error())
		return ctgry, internal.ErrFailedToInsertCategory
	}

	// Get the last inserted ID
	categoryID, err := res.LastInsertId()
	if err != nil {
		log.Println("error occurred in fetching the inserted ID:", err.Error())
		return ctgry, internal.ErrFailedToInsertCategory
	}

	// Assign values to the category object
	ctgry.ID = int(categoryID)
	ctgry.Name = category.Name
	ctgry.Description = category.Description
	ctgry.IsActive = category.IsActive

	return ctgry, nil
}

func (cts *categoryStore) UpdateCategory(ctx context.Context, category dto.Category) (dto.Category, error) {
	// Check if the category ID exists
	if !cts.categoryExists(category.ID) {
		log.Println("category with ID:", category.ID, "does not exist")
		return category, internal.ErrCategoryNotFound
	}

	// Check if the updated name already exists
	if cts.categoryNameExists(category.ID, category.Name) {
		log.Println("category with name:", category.Name, "already exists")
		return category, internal.ErrCategoryNameExists
	}

	// Prepare the update statement
	query := "UPDATE category SET name = ?, description = ?, is_active = ? WHERE id = ?"
	statement, err := cts.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occurred in preparing the update statement:", err.Error())
		return category, internal.ErrFailedToUpdateCategory
	}
	defer statement.Close()

	// Execute the update statement
	res, err := statement.Exec(category.Name, category.Description, category.IsActive, category.ID)
	if err != nil {
		log.Println("error occurred in executing the update statement:", err.Error())
		return category, internal.ErrFailedToUpdateCategory
	}

	// Check if any rows were affected
	noOfRowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("error occurred in fetching the number of affected rows:", err.Error())
		return category, internal.ErrFailedToUpdateCategory
	}

	// If no rows were affected, return an error
	if noOfRowsAffected == 0 {
		log.Println("no rows were affected by the update operation")
		return category, internal.ErrCategoryNotFound
	}

	return category, nil
}

// Function to check if the category with the specified ID exists
func (cts *categoryStore) categoryExists(categoryID int) bool {
	query := "SELECT COUNT(*) FROM category WHERE id = ?"
	var count int
	err := cts.BaseRepsitory.DB.QueryRow(query, categoryID).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking category existence:", err.Error())
		return false
	}
	return count > 0
}

// Function to check if the updated name already exists
func (cts *categoryStore) categoryNameExists(categoryID int, updatedName string) bool {
	query := "SELECT COUNT(*) FROM category WHERE name = ? AND id != ?"
	var count int
	err := cts.BaseRepsitory.DB.QueryRow(query, updatedName, categoryID).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking category name existence:", err.Error())
		return false
	}
	return count > 0
}
