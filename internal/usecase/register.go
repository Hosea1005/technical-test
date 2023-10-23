package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"service-auth/internal/http/register/request"
	"service-auth/internal/http/register/response"
)

func (a AuthUseCase) Register(ctx context.Context, request request.RegisterRequest) *response.RegisterResponse {
	// hash password using bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	request.Password = string(hashPassword)
	user, err := a.UserRepository.InsertUser(ctx, request)
	if err != nil {
		return &response.RegisterResponse{
			Status: response.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	errWall := a.UserRepository.InsertWallet(ctx, int64(user.ID))
	if errWall != nil {
		return &response.RegisterResponse{
			Status: response.Status{
				Code:    400,
				Message: errWall.Error(),
			},
		}
	}

	return &response.RegisterResponse{
		Status: response.Status{
			Code:    200,
			Message: "Success Register",
		},
	}
}
