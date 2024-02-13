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
	Signup(ctx context.Context, user dto.UserSignUpRequest) (dto.UserLoginResponse, error)
	Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error)
	GetUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUser(ctx context.Context, userID int) (dto.UserResponse, error)
	UpdateUser(ctx context.Context, updateInfo dto.UpdateUserInfo, userID int) (dto.UserResponse, error)
}
type service struct {
	userRepo repository.UserStorer
}

func NewService(userRepo repository.UserStorer) Service {
	return &service{
		userRepo: userRepo,
	}
}
func (sgSrv *service) Signup(ctx context.Context, user dto.UserSignUpRequest) (dto.UserLoginResponse, error) {

	valUser := validateUser(&user)
	if !valUser {
		return dto.UserLoginResponse{}, errors.New("enter valid data")
	}
	user.Password = HashPassword(user.Password)
	resp, err := sgSrv.userRepo.Signup(ctx, user)
	return resp, err
}
func (sgSrv *service) Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error) {

	valEmailPassword := isValidEmail(user.Email) && isValidPassword(user.Password)
	if !valEmailPassword {
		return dto.UserLoginResponse{}, errors.New("enter valid email and password")
	}
	// user.Password = HashPassword(user.Password)
	// userData := dto.UserLoginResponse{}
	userData, err := sgSrv.userRepo.Login(ctx, user)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	log.Printf("login service response: %v", userData)

	return userData, err

}
func (sgSrv *service) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {

	userList, err := sgSrv.userRepo.GetUsers(ctx)
	if err != nil {
		return userList, err
	}
	return userList, nil
}
func (sgSrv *service) GetUser(ctx context.Context, userID int) (dto.UserResponse, error) {
	if userID <= 0 {
		return dto.UserResponse{}, errors.New("invalid user id")
	}
	user, err := sgSrv.userRepo.GetUser(ctx, userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (sgSrv *service) UpdateUser(ctx context.Context, updateInfo dto.UpdateUserInfo, userID int) (dto.UserResponse, error) {
	updateInfo.Password = HashPassword(updateInfo.Password)
	userResp, err := sgSrv.userRepo.UpdateUser(ctx, updateInfo, userID)
	if err != nil {
		log.Println(err.Error())
		return dto.UserResponse{}, err
	}
	return userResp, nil
}
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}
