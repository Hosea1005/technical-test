package usecase

import (
	"context"
	dRequest "service-auth/internal/http/detail_player/request"
	dResponse "service-auth/internal/http/detail_player/response"
	"strconv"
)

func (a AuthUseCase) GetDetailPlayer(ctx context.Context, request dRequest.DetailPlayerRequest) *dResponse.DetailPlayerResponse {
	_, errToken := a.CacheService.GetToken(ctx, request.Token)
	if errToken != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: "Unauthorized",
			},
		}
	}
	id, err := strconv.ParseInt(request.ID, 0, 64)
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resUser, err := a.UserRepository.GetDetailPlayer(ctx, id)
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resAccount, err := a.UserRepository.GetAccount(ctx, int64(resUser.ID))
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resAccount, err = a.UserRepository.GetAccount(ctx, int64(resUser.ID))
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resBank, err := a.UserRepository.GetBank(ctx, resAccount.IDBank)
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	resWallet, err := a.UserRepository.GetWallet(ctx, int64(resUser.ID))
	if err != nil {
		return &dResponse.DetailPlayerResponse{
			Status: dResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}

	return &dResponse.DetailPlayerResponse{
		Status: dResponse.Status{
			Code:    200,
			Message: "Success",
		},
		Data: dResponse.Data{
			Player: dResponse.Player{
				ID:        resUser.ID,
				Name:      resUser.Fullname,
				Username:  resUser.Username,
				Level:     resUser.Level,
				CreatedAt: resUser.CreatedAt,
				Account: dResponse.Account{
					ID:      resAccount.ID,
					Grade:   resAccount.Grade,
					Balance: resAccount.Balance,
					Bank:    resBank.Name,
				},
				Wallet: dResponse.Wallet{
					ID:        resWallet.ID,
					Balance:   resWallet.Balance,
					CreatedAt: resWallet.CreatedAt,
				},
			},
		},
	}
}
