package repository

import (
	"context"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
)

type UserStorer interface {
	Signup(ctx context.Context, user dto.UserSignUpRequest) (dto.UserLoginResponse, error)
	Login(ctx context.Context, user dto.UserLoginRequest) (dto.UserLoginResponse, error)
	GetUsers(ctx context.Context) ([]dto.UserResponse, error)
	GetUser(ctx context.Context, userId int) (dto.UserResponse, error)
	UpdateUser(ctx context.Context,updateInfo dto.UpdateUserInfo,userID int)(dto.UserResponse,error)
}
type User struct {
	ID        int
	Phone     string
	Email     string
	Password  string
	Firstname string
	Lastname  string
	Role      string
}
