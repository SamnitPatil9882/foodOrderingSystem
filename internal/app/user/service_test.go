package user

import (
	"context"
	"testing"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	type loginDetails struct {
		lgDet dto.UserLoginRequest
	}
	test := []struct {
		name     string
		args     loginDetails
		wantErr  bool
		wantResp dto.UserLoginResponse
		setup    func(userMock *mocks.UserStorer)
	}{
		{
			name: "successful",
			args: loginDetails{
				dto.UserLoginRequest{
					Email:    "patil@gmail.com",
					Password: "patil@123",
				},
			},
			wantErr: true,
			wantResp: dto.UserLoginResponse{
				ID:   1,
				Role: "Customer",
			},
			setup: func(userMock *mocks.UserStorer) {
				userMock.On("Login", mock.Anything, mock.Anything).Return(dto.UserResponse{ID: 1,Role: "customer"}, nil).Once()
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			userstore := mocks.NewUserStorer(t)
			svc := NewService(userstore)

			svc.Login(context.Background(), tt.args.lgDet)

			_, err := svc.Login(context.Background(), tt.args.lgDet)
			if (err != nil) != tt.wantErr {
				t.Errorf("errr %s", err)
			}
		})
	}

}
