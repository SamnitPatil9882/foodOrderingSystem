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
		fmt.Println("error occured in selecting food : " + err.Error())
		return foodList, err
	}
	defer rows.Close()
	for rows.Next() {
		food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		fmt.Println(food)
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
	defer rows.Close()
	for rows.Next() {
		food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		foodList = append(foodList, food)
	}
	return foodList, nil
}

func (fds *FoodStore) CreateFood(ctx context.Context, fd dto.FoodCreateRequest) (repository.Food, error) {
	food := repository.Food{}
	query := "INSERT INTO food (name,price,category_id,is_veg,is_avail) VALUES(?,?,?,?,?)"
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return food, err
	}
	defer statement.Close()
	res, err := statement.Exec(fd.Name, fd.Price, fd.CategoryID, fd.IsVeg, fd.IsAvail)
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
	defer rows.Close()
	for rows.Next() {
		// food := repository.Food{}
		rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		// foodList = append(foodList, food)
	}
	fmt.Println(food)
	return food, nil
}

func (fds *FoodStore) UpdateFood(ctx context.Context, food dto.Food) (repository.Food, error) {
	resFood := repository.Food{}

	query := fmt.Sprintf("UPDATE food SET category_id = %d,price=%d,name=\"%s\",is_veg=%d ,is_avail = %d WHERE id=%d", food.CategoryID, food.Price, food.Name, food.IsVeg, food.IsAvail, food.ID)
	statement, err := fds.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error occured in update prepare statement: " + err.Error())
		return resFood, err
	}
	defer statement.Close()
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
	defer rows.Close()
	for rows.Next() {
		// food := repository.Food{}
		rows.Scan(&resFood.ID, &resFood.CategoryID, &resFood.Price, &resFood.Name, &resFood.IsVeg, &resFood.IsAvail)
		// foodList = append(foodList, food)
	}
	// fmt.Println(food)
	return resFood, nil
}

func (fds *FoodStore) GetFoodByID(ctx context.Context, FoodID int64) (repository.Food, error) {
	food := repository.Food{}
	fmt.Printf("in db %d", FoodID)
	// query := fmt.Sprintf("SELECT * FROM food WHERE id = %d", FoodID)
	query := fmt.Sprintf(`
	SELECT f.id, f.category_id, f.price, f.name, f.is_veg, f.is_avail
	FROM food f
	JOIN category c ON f.category_id = c.id
	WHERE f.id = %d AND c.is_active = 1 AND f.is_avail = 1
	`, FoodID)

	rows, err := fds.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("error occured in selecting food by id: " + err.Error())
		return food, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&food.ID, &food.CategoryID, &food.Price, &food.Name, &food.IsVeg, &food.IsAvail)
		return food, err
	}
	return food, errors.New("no match found")
}
