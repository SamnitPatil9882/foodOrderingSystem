package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail); err != nil {
			return foodList, fmt.Errorf("failed to scan food row: %v", err)
		}
		foodList = append(foodList, food)
	}
	return foodList, nil
}

func (fds *FoodStore) GetFoodByCategory(ctx context.Context, categoryID int) ([]repository.Food, error) {
	foodList := make([]repository.Food, 0)
	query := fmt.Sprintf("SELECT * FROM food WHERE category_id=%d", categoryID)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return foodList, fmt.Errorf("failed to fetch food by category: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		food := repository.Food{}
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail); err != nil {
			return foodList, fmt.Errorf("failed to scan food row: %v", err)
		}
		foodList = append(foodList, food)
	}
	return foodList, nil
}

func (fds *FoodStore) CreateFood(ctx context.Context, fd dto.FoodCreateRequest) (repository.Food, error) {
	food := repository.Food{}
	query := "INSERT INTO food (name,price,category_id,is_veg,is_avail) VALUES(?,?,?,?,?)"
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		return food, fmt.Errorf("failed to prepare food insertion: %v", err)
	}
	defer statement.Close()
	res, err := statement.Exec(fd.Name, fd.Price, fd.CategoryID, fd.IsVeg, fd.IsAvail)
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
	for rows.Next() {
		if err := rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail); err != nil {
			return food, fmt.Errorf("failed to scan food row: %v", err)
		}
	}
	return food, nil
}

func (fds *FoodStore) UpdateFood(ctx context.Context, food dto.Food) (repository.Food, error) {
	resFood := repository.Food{}

	query := fmt.Sprintf("UPDATE food SET category_id = %d,price=%d,name=\"%s\",is_veg=%d ,is_avail = %d WHERE id=%d", food.CategoryID, food.Price, food.Name, food.IsVeg, food.IsAvail, food.ID)
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		return resFood, fmt.Errorf("failed to prepare food update: %v", err)
	}
	defer statement.Close()
	res, err := statement.Exec()
	if err != nil {
		return resFood, fmt.Errorf("failed to execute food update query: %v", err)
	}
	noOfRawAffected, err := res.RowsAffected()
	if err != nil {
		return resFood, fmt.Errorf("failed to get rows affected: %v", err)
	}
	if noOfRawAffected == 0 {
		return resFood, errors.New("no rows affected")
	}

	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", food.ID)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return resFood, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&resFood.ID, &resFood.CategoryID, &resFood.Price, &resFood.Name, &resFood.IsVeg, &resFood.IsAvail); err != nil {
			return resFood, fmt.Errorf("failed to scan food row: %v", err)
		}
	}
	return resFood, nil
}

func (fds *FoodStore) GetFoodByID(ctx context.Context, FoodID int64) (repository.Food, error) {
	food := repository.Food{}

	query := fmt.Sprintf(`
	SELECT f.id, f.category_id, f.price, f.name, f.is_veg, f.is_avail
	FROM food f
	JOIN category c ON f.category_id = c.id
	WHERE f.id = %d AND c.is_active = 1 AND f.is_avail = 1
	`, FoodID)

	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return food, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		return food, err
	}
	return food, errors.New("no match found")
}

func (fds *FoodStore) GetFoodInfoByID(ctx context.Context, FoodID int64) (dto.Food, error) {
	food := dto.Food{}

	query := fmt.Sprintf("SELECT * FROM food WHERE id = %d", FoodID)

	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		return food, fmt.Errorf("failed to fetch food by ID: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		return food, err
	}
	return food, errors.New("no match found")
}
