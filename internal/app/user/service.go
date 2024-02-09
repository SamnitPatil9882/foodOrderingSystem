package user

import (
	"context"
	"errors"
	"log"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Signup(ctx context.Context, user dto.UserSignUpRequest) error
	Login(ctx context.Context, user dto.UserLoginRequest) error
	GetUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUser(ctx context.Context, userID int) (dto.UserResponse, error)
}
type service struct {
	userRepo repository.UserStorer
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}
func (sgSrv *service) Signup(ctx context.Context, user dto.UserSignUpRequest) error {

	valUser := validateUser(user)
	if !valUser {
		return errors.New("enter valid data")
	}
	user.Password = HashPassword(user.Password)
	err := sgSrv.userRepo.Signup(ctx, user)
	return err
}
func (sgSrv *service) Login(ctx context.Context, user dto.UserLoginRequest) error {

	valEmailPassword := isValidEmail(user.Email) && isValidPassword(user.Password)
	if !valEmailPassword {
		return errors.New("enter valid email and password")
	}
	// user.Password = HashPassword(user.Password)
	err := sgSrv.userRepo.Login(ctx, user)
	if err != nil {
		return err
	}
	return nil

}
func (sgSrv *service) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {

	userList, err := sgSrv.userRepo.GetUsers(ctx)
	if err != nil {
		return userList, err
	}
	return userList, nil
}
func (sgSrv *service) GetUser(ctx context.Context, userID int) (dto.UserResponse, error) {

	user, err := sgSrv.userRepo.GetUser(ctx, userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}
