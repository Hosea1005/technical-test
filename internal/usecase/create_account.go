package usecase

import (
	"context"
	"service-auth/helper"
	request2 "service-auth/internal/http/create_account/request"
	response2 "service-auth/internal/http/create_account/response"
)

func (a AuthUseCase) CreateAccount(ctx context.Context, request request2.CreateAccountRequest) *response2.CreateAccountResponse {
	_, errToken := a.CacheService.GetToken(ctx, request.Token)
	if errToken != nil {
		return &response2.CreateAccountResponse{
			Status: response2.Status{
				Code:    400,
				Message: "Unauthorized",
			},
		}
	}
	dataToken, err := helper.ParseJWTToken(request.Token)
	if err != nil {
		return &response2.CreateAccountResponse{
			Status: response2.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	errBank := a.UserRepository.CheckIdBank(ctx, request.IDBank)
	if errBank != nil {
		return &response2.CreateAccountResponse{
			Status: response2.Status{
				Code:    400,
				Message: errBank.Error(),
			},
		}
	}
	request.IDUser = dataToken.UserId
	err = a.UserRepository.InsertAccount(ctx, request)
	if err != nil {
		return &response2.CreateAccountResponse{
			Status: response2.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &response2.CreateAccountResponse{
		Status: response2.Status{
			Code:    200,
			Message: "Success Create Account Bank",
		},
	}
}
