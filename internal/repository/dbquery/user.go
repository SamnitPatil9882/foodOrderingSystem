package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
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

func (us *UserStore) Signup(ctx context.Context, user dto.UserSignUpRequest) (dto.UserLoginResponse, error) {
    // Check if email already exists
    emailExistsQuery := "SELECT COUNT(*) FROM user WHERE email = ?"
    var emailCount int
    err := us.BaseRepsitory.DB.QueryRow(emailExistsQuery, user.Email).Scan(&emailCount)
    if err != nil {
        return dto.UserLoginResponse{}, err
    }
    if emailCount > 0 {
        return dto.UserLoginResponse{}, internal.ErrEmailExists
    }

    // Check if phone number already exists
    phoneExistsQuery := "SELECT COUNT(*) FROM user WHERE phone = ?"
    var phoneCount int
    err = us.BaseRepsitory.DB.QueryRow(phoneExistsQuery, user.Phone).Scan(&phoneCount)
    if err != nil {
        return dto.UserLoginResponse{}, err
    }
    if phoneCount > 0 {
        return dto.UserLoginResponse{}, internal.ErrPhoneExists
    }

    // Insert new user
    query := "INSERT INTO user (phone,email,password,firstname,lastname,role) VALUES(?,?,?,?,?,?)"
    statement, err := us.BaseRepsitory.DB.Prepare(query)
    if err != nil {
        fmt.Println("error in inserting: " + err.Error())
        return dto.UserLoginResponse{}, err
    }
    defer statement.Close()

    res, err := statement.Exec(user.Phone, user.Email, user.Password, user.Firstname, user.Lastname, user.Role)
    if err != nil {
        fmt.Println("error occured in executing insert query: " + err.Error())
        return dto.UserLoginResponse{}, errors.New(internal.InternalServerError)
    }
    id, err := res.LastInsertId()
    if err != nil {
        fmt.Println("error in executing: " + err.Error())
        return dto.UserLoginResponse{}, errors.New(internal.UserAvail)
    }
    resp := dto.UserLoginResponse{ID: int(id), Role: user.Role}
    return resp, nil
}

func (us *UserStore) Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error) {
    response := dto.UserLoginResponse{}
    // Check if user with provided email exists
    query := fmt.Sprintf("SELECT password,id,role FROM user WHERE email = \"%s\"", user.Email)
    rows, err := us.BaseRepsitory.DB.Query(query)
    if err != nil {
        fmt.Println("Email is incorrect: " + err.Error())
        return response, internal.ErrEmailOrPasswordIncorrect
    }
    defer rows.Close()

    var password string
    var id int
    var role string

    for rows.Next() {
        err := rows.Scan(&password, &id, &role)
        if err != nil {
            return dto.UserLoginResponse{}, err
        }
    }
    if id == 0 {
        return dto.UserLoginResponse{}, internal.ErrEmailOrPasswordIncorrect
    }

    err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
    if err != nil {
        return dto.UserLoginResponse{}, internal.ErrEmailOrPasswordIncorrect
    }
    response.ID = id
    response.Role = role
    return response, nil
}


func (us *UserStore) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {
	usersList := make([]dto.UserResponse, 0)
	query := "SELECT id,firstname,lastname,email,phone,role FROM user "
	rows, err := us.BaseRepsitory.DB.Query(query)
	if err != nil {
		return usersList, err
	}
	defer rows.Close()
	for rows.Next() {
		user := dto.UserResponse{}
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Role)
		if err != nil {
			return []dto.UserResponse{}, err
		}
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
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Role)
		return user, err
	}
	return user, internal.ErrUserNotFound
}
func (us *UserStore) UpdateUser(ctx context.Context, updateInfo dto.UpdateUserInfo, userID int) (dto.UserResponse, error) {
	log.Printf("%d   %s", userID, updateInfo.Email)

	// Prepare the SQL statement with placeholders
	_,err:=us.GetUser(ctx,userID)
	if err!= nil {
        return dto.UserResponse{}, internal.ErrUserNotFound;
    }
	query := `UPDATE user
	SET phone = ?, email = ?, password = ?, firstname = ?, lastname = ?
	WHERE id = ?`

	// Prepare the statement
	statement, err := us.BaseRepsitory.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return dto.UserResponse{}, errors.New(internal.InternalServerError)
	}
	defer statement.Close()


	// Execute the statement with the actual values
	res, err := statement.ExecContext(ctx, updateInfo.Phone, updateInfo.Email, updateInfo.Password, updateInfo.Firstname, updateInfo.Lastname, userID)
	if err != nil {
		log.Println(err.Error())
		return dto.UserResponse{}, errors.New(internal.InternalServerError)
	}

	noOfrowAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return dto.UserResponse{}, errors.New(internal.InternalServerError)
	}
	if noOfrowAffected == 0 {
		log.Println("No rows affected")
		return dto.UserResponse{}, errors.New(internal.InternalServerError)
	}
	userInfo, err := us.GetUser(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		return dto.UserResponse{}, errors.New(internal.InternalServerError)
	}
	return userInfo, nil
}
