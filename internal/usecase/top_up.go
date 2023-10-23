package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"service-auth/helper"
	request3 "service-auth/internal/http/top_up/request"
	response3 "service-auth/internal/http/top_up/response"
	"service-auth/internal/model"
	"time"
)

func (a AuthUseCase) TopUp(ctx context.Context, request request3.TopUpRequest) *response3.TopUpResponse {
	_, errToken := a.CacheService.GetToken(ctx, request.Token)
	if errToken != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: "Unauthorized",
			},
		}
	}
	dataToken, err := helper.ParseJWTToken(request.Token)
	if err != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    500,
				Message: err.Error(),
			},
		}
	}
	request.IDUser = dataToken.UserId
	var total, balance float64
	resAcc, err := a.UserRepository.CheckAccount(ctx, request.IDUser)
	if err != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	if resAcc.Balance < request.Total {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: "your balance is not enough",
			},
		}
	}

	total = resAcc.Balance - request.Total

	err = a.UserRepository.UpdateBalance(ctx, int64(resAcc.ID), total)
	if err != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resWall, errWallet := a.UserRepository.GetWallet(ctx, request.IDUser)
	if errWallet != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: errWallet.Error(),
			},
		}
	}
	balance = resWall.Balance + request.Total

	err = a.UserRepository.UpdateWallet(ctx, int64(resWall.ID), balance)
	if err != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100000000)

	randomString := fmt.Sprintf("%08d", randomNumber)

	dataCache := model.TopUp{
		IDUser:    request.IDUser,
		Total:     request.Total,
		IDAccount: int64(resAcc.ID),
		IDWallet:  int64(resWall.ID),
		CreatedAt: time.Now(),
	}
	err = a.CacheService.Set(ctx, randomString, &dataCache)
	if err != nil {
		return &response3.TopUpResponse{
			Status: response3.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	return &response3.TopUpResponse{
		Status: response3.Status{
			Code:    200,
			Message: "Success Top-Up",
		},
	}
}
