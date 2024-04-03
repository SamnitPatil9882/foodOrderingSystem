package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type FoodStore struct {
	BaseRepsitory
}

func NewFoodRepo(db *sql.DB) repository.FoodStorer {
	return &FoodStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (fds *FoodStore) GetListOfOrder(ctx context.Context) ([]repository.Food, error) {
	foodList := make([]repository.Food, 0)
	query := "SELECT * FROM food ORDER BY category_id"
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return foodList, fmt.Errorf("failed to fetch list of food: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		food := repository.Food{}
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.Description, &food.ImgUrl, &food.IsVeg, &food.IsAvail); err != nil {
			return foodList, fmt.Errorf("failed to scan food row: %v", err)
		}
		foodList = append(foodList, food)
	}
	return foodList, nil
}

func (fds *FoodStore) GetFoodByCategory(ctx context.Context, categoryID int) ([]repository.Food, error) {
	// Check if the category ID exists
	if !fds.categoryExists(categoryID) {
		return nil, internal.ErrCategoryNotFound
	}

	foodList := make([]repository.Food, 0)
	query := fmt.Sprintf("SELECT * FROM food WHERE category_id=%d", categoryID)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return foodList, fmt.Errorf("failed to fetch food by category: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		food := repository.Food{}
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.Description, &food.ImgUrl, &food.IsVeg, &food.IsAvail); err != nil {
			return foodList, fmt.Errorf("failed to scan food row: %v", err)
		}
		foodList = append(foodList, food)
	}
	return foodList, nil
}

// Function to check if the category with the specified ID exists
func (fds *FoodStore) categoryExists(categoryID int) bool {
	query := "SELECT COUNT(*) FROM category WHERE id = ?"
	var count int
	err := fds.BaseRepsitory.DB.QueryRow(query, categoryID).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking category existence:", err.Error())
		return false
	}
	return count > 0
}

func (fds *FoodStore) CreateFood(ctx context.Context, fd dto.FoodCreateRequest) (repository.Food, error) {
	food := repository.Food{}

    log.Println("food creation : dto FoodCreateRequest: ", fd)

	// Check if the food name already exists
	if fds.foodNameExists(fd.Name) {
		return food, internal.ErrFoodNameExists
	}

	// Check if the food category exists
	if !fds.categoryExists(int(fd.CategoryID)) {
		return food, internal.ErrCategoryNotFound
	}

	query := "INSERT INTO food (name,price,category_id,is_veg,is_avail,description,imgurl) VALUES(?,?,?,?,?,?,?)"
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		return food, fmt.Errorf("failed to prepare food insertion: %v", err)
	}
	defer statement.Close()

	res, err := statement.Exec(fd.Name, fd.Price, fd.CategoryID, fd.IsVeg, fd.IsAvail,fd.Description,fd.ImgUrl)
	if err != nil {
		return food, fmt.Errorf("failed to execute food insertion query: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return food, fmt.Errorf("failed to fetch last insert ID: %v", err)
	}

	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", id)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return food, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()
    log.Println(("row next"))
    for rows.Next() {
        if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.Description, &food.ImgUrl, &food.IsVeg, &food.IsAvail); err != nil {
            log.Println("food: ",food)
            return food, fmt.Errorf("failed to scan food row: %v", err)
        }
    }
	return food, nil
}

// Function to check if the food name already exists
func (fds *FoodStore) foodNameExists(foodName string) bool {
	query := "SELECT COUNT(*) FROM food WHERE name = ?"
	var count int
	err := fds.BaseRepsitory.DB.QueryRow(query, foodName).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking food name existence:", err.Error())
		return false
	}
	return count > 0
}

// func (fds *FoodStore) UpdateFood(ctx context.Context, food dto.Food) (repository.Food, error) {
// 	resFood := repository.Food{}

// 	// Check if the food category exists
// 	if !fds.categoryExists(int(food.CategoryID)) {
// 		return resFood, internal.ErrCategoryNotFound
// 	}

// 	// Check if the food name already exists for other food items except the one being updated
// 	if fds.foodNameExistsOtherThanCurrent(food.Name, food.ID) {
// 		return resFood, internal.ErrFoodNameExists
// 	}

//     query := "UPDATE food SET category_id = ?, price = ?, name = ?, description = ?, imgurl = ?, is_veg = ?, is_avail = ? WHERE id = ?"
//     statement, err := fds.DB.Prepare(query)
//     if err != nil {
//         return resFood, fmt.Errorf("failed to prepare food update: %v", err)
//     }
//     defer statement.Close()

// 	res, err := statement.Exec(food.CategoryID,food.Price,food.Name,food.Description,food.ImgUrl,food.IsVeg,food.IsAvail,food.ID)
// 	if err != nil {
// 		return resFood, fmt.Errorf("failed to execute food update query: %v", err)
// 	}

// 	noOfRawAffected, err := res.RowsAffected()
// 	if err != nil {
// 		return resFood, fmt.Errorf("failed to get rows affected: %v", err)
// 	}
// 	if noOfRawAffected == 0 {
// 		return resFood, errors.New("no rows affected")
// 	}

// 	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", food.ID)
// 	rows, err := fds.BaseRepsitory.DB.Query(query)
// 	if err != nil {
// 		return resFood, fmt.Errorf("failed to fetch food by ID: %v", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		if err := rows.Scan(&resFood.ID, &resFood.CategoryID, &resFood.Price,&resFood.Description,&resFood.ImgUrl ,&resFood.Name, &resFood.IsVeg, &resFood.IsAvail); err != nil {
// 			return resFood, fmt.Errorf("failed to scan food row: %v", err)
// 		}
// 	}
// 	return resFood, nil
// }

func (fds *FoodStore) UpdateFood(ctx context.Context, food dto.Food) (repository.Food, error) {
	resFood := repository.Food{}

	// Check if the food category exists


	// Check if the food name already exists for other food items except the one being updated


	// Initialize an empty list to hold the column-value pairs for the update query
	var updates []string
	var args []interface{}

	// Append each non-zero field to the updates list along with its corresponding argument
	if food.CategoryID != -1 {
		if !fds.categoryExists(int(food.CategoryID)) {
			return resFood, internal.ErrCategoryNotFound
		}
		updates = append(updates, "category_id = ?")
		args = append(args, food.CategoryID)
	}
	if food.Price != -1 {
		updates = append(updates, "price = ?")
		args = append(args, food.Price)
	}
	if food.Name != "" {
		if fds.foodNameExistsOtherThanCurrent(food.Name, food.ID) {
			return resFood, internal.ErrFoodNameExists
		}
		updates = append(updates, "name = ?")
		args = append(args, food.Name)
	}
	if food.Description != "" {
		updates = append(updates, "description = ?")
		args = append(args, food.Description)
	}
	if food.ImgUrl != "" {
		updates = append(updates, "imgurl = ?")
		args = append(args, food.ImgUrl)
	}
	if food.IsAvail != -1 {
		updates = append(updates, "is_avail = ?")
		args = append(args, food.IsAvail)
	}
	// Similarly, add conditions for other fields as needed

	// Combine all update statements into a single query
	query := fmt.Sprintf("UPDATE food SET %s WHERE id = ?", strings.Join(updates, ", "))

	// Prepare the update statement
	statement, err := fds.DB.Prepare(query)
	if err != nil {
		return resFood, fmt.Errorf("failed to prepare food update: %v", err)
	}
	defer statement.Close()

	// Execute the update statement with the provided arguments
	args = append(args, food.ID)
	res, err := statement.Exec(args...)
	if err != nil {
		return resFood, fmt.Errorf("failed to execute food update query: %v", err)
	}

	// Check the number of affected rows
	noOfRawAffected, err := res.RowsAffected()
	if err != nil {
		return resFood, fmt.Errorf("failed to get rows affected: %v", err)
	}
	if noOfRawAffected == 0 {
		return resFood, errors.New("no rows affected")
	}

	// Fetch the updated food item from the database
	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", food.ID)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return resFood, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()

	// Scan the fetched row into the result struct
	for rows.Next() {
		if err := rows.Scan(&resFood.ID, &resFood.CategoryID, &resFood.Price, &resFood.Description, &resFood.ImgUrl, &resFood.Name, &resFood.IsVeg, &resFood.IsAvail); err != nil {
			return resFood, fmt.Errorf("failed to scan food row: %v", err)
		}
	}
	return resFood, nil
}


// Function to check if the food name already exists for other food items except the one being updated
func (fds *FoodStore) foodNameExistsOtherThanCurrent(foodName string, foodID int64) bool {
	query := "SELECT COUNT(*) FROM food WHERE name = ? AND id != ?"
	var count int
	err := fds.BaseRepsitory.DB.QueryRow(query, foodName, foodID).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking food name existence:", err.Error())
		return false
	}
	return count > 0
}

// Function to check if the food name already exists for other food items except the one being updated

func (fds *FoodStore) GetFoodByID(ctx context.Context, foodID int64) (repository.Food, error) {
	food := repository.Food{}

	// Check if food ID exists
	if !fds.foodExists(foodID) {
		return food, internal.ErrFoodNotFound
	}

	query := fmt.Sprintf(`
        SELECT f.id, f.category_id, f.price, f.name,f.description, f.imgurl, f.is_veg, f.is_avail
        FROM food f
        JOIN category c ON f.category_id = c.id
        WHERE f.id = %d AND c.is_active = 1 AND f.is_avail = 1
    `, foodID)

	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return food, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()

	
	for rows.Next() {
        err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.Description, &food.ImgUrl, &food.IsVeg, &food.IsAvail)
		if err != nil {
			return food, fmt.Errorf("failed to scan food row: %v", err)
		}
		return food, nil
	}
	fmt.Println("food by id: ",food)
	return food, errors.New("no match found")
}

// Function to check if the food ID exists
func (fds *FoodStore) foodExists(foodID int64) bool {
	query := "SELECT COUNT(*) FROM food WHERE id = ?"
	var count int
	err := fds.BaseRepsitory.DB.QueryRow(query, foodID).Scan(&count)
	if err != nil {
		log.Println("error occurred in checking food existence:", err.Error())
		return false
	}
	return count > 0
}

func (fds *FoodStore) GetFoodInfoByID(ctx context.Context, foodID int64) (dto.Food, error) {
	food := dto.Food{}

	// Check if food ID exists
	if !fds.foodExists(foodID) {
		return food, internal.ErrFoodNotFound
	}

	query := fmt.Sprintf("SELECT * FROM food WHERE id = %d", foodID)

	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return food, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price,&food.Name , &food.Description, &food.ImgUrl, &food.IsVeg, &food.IsAvail); err != nil {
			return food, fmt.Errorf("failed to scan food row: %v", err)
		}
		fmt.Println("food info by id: ",food);
		return food, nil
	}

	return food, errors.New("no match found")
}
