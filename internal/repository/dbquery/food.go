package dbquery

import (
	"context"
	"database/sql"
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
		fmt.Println("error occured in selecting food : " + err.Error())
		return foodList, err
	}

	for rows.Next() {
		food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg)
		foodList = append(foodList, food)
	}
	return foodList, nil
}
func (fds *FoodStore) GetFoodByCategory(ctx context.Context, categoryID int) ([]repository.Food, error) {
	foodList := make([]repository.Food, 0)
	query := fmt.Sprintf("SELECT * FROM food WHERE category_id=%d", categoryID)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting food by category: " + err.Error())
	}

	for rows.Next() {
		food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg)
		foodList = append(foodList, food)
	}
	return foodList, nil
}

func (fds *FoodStore) CreateFood(ctx context.Context, fd dto.FoodCreateRequest) (repository.Food, error) {
	food := repository.Food{}
	query := "INSERT INTO food (name,price,category_id,is_veg) VALUES(?,?,?,?)"
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return food, err
	}
	res, err := statement.Exec(fd.Name, fd.Price, fd.CategoryID, fd.IsVeg)
	if err != nil {
		fmt.Println("error occured in executing insert query: " + err.Error())
		return food, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error occured in getting lastInsertId: " + err.Error())
		return food, err
	}

	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", id)
	fmt.Println(id)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting food by id: " + err.Error())
		return food, err
	}

	for rows.Next() {
		// food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg)
		// foodList = append(foodList, food)
	}
	fmt.Println(food)
	return food, nil
}

func (fds *FoodStore) UpdateFood(ctx context.Context, food dto.Food) (repository.Food, error) {
	resFood := repository.Food{}

	query := fmt.Sprintf("UPDATE food SET category_id = %d,price=%d,name=\"%s\",is_veg=%d WHERE id=%d", food.CategoryID, food.Price, food.Name, food.IsVeg, food.ID)
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error occured in update prepare statement: " + err.Error())
		return resFood, err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Println("error occured in execution of update food query: " + err.Error())
		return resFood, err
	}

	query = fmt.Sprintf("SELECT * FROM food WHERE id=%d", food.ID)
	// fmt.Println(id)
	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting food by id: " + err.Error())
		return resFood, err
	}

	for rows.Next() {
		// food := repository.Food{}
		rows.Scan(&resFood.ID, &resFood.CategoryID, &resFood.Price, &resFood.Name, &resFood.IsVeg)
		// foodList = append(foodList, food)
	}
	// fmt.Println(food)
	return resFood, nil
}
