package dbquery

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	BaseRepsitory
}

func NewUserRepo(db *sql.DB) repository.UserStorer {
	return &UserStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (us *UserStore) Signup(ctx context.Context, user dto.UserSignUpRequest) error {

	query := "INSERT INTO user (phone,email,password,firstname,lastname,role) VALUES(?,?,?,?,?,?)"
	statement, err := us.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return err
	}
	_, err = statement.Exec(user.Phone, user.Email, user.Password, user.Firstname, user.Lastname, user.Role)
	if err != nil {
		fmt.Println("error occured in executing insert query: " + err.Error())
		return err
	}

	return nil
}

func (us *UserStore) Login(ctx context.Context, user dto.UserLoginRequest) error {

	query := fmt.Sprintf("SELECT password FROM user WHERE email = \"%s\"", user.Email)
	rows, err := us.BaseRepsitory.DB.Query(query)
	if err != nil {
		fmt.Println("Email is incorrect: " + err.Error())
		return err
	}
	var password string
	for rows.Next() {
		// food := repository.Food{}
		rows.Scan(&password)
		// foodList = append(foodList, food)
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
	// if user.Password == password {
	// 	return nil
	// }
	// return errors.New("password is incorrect")

}

func (us *UserStore) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {
	usersList := make([]dto.UserResponse, 0)
	query := "SELECT id,firstname,lastname,email,phone,role FROM user "
	rows, err := us.BaseRepsitory.DB.Query(query)
	if err != nil {
		return usersList, err
	}
	for rows.Next() {
		user := dto.UserResponse{}
		rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Role)
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (us *UserStore) GetUser(ctx context.Context, userId int) (dto.UserResponse, error) {
	user := dto.UserResponse{}

	query := fmt.Sprintf("SELECT id,firstname,lastname,email,phone,role FROM user WHERE id=%d", userId)
	rows, err := us.BaseRepsitory.DB.Query(query)
	if err != nil {
		return user, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Role)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
