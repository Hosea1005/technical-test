package domain

import (
	"context"
	request2 "service-auth/internal/http/create_account/request"
	"service-auth/internal/http/register/request"
	"service-auth/internal/model"
)

type AuthRepository interface {
	InsertUser(ctx context.Context, data request.RegisterRequest) (*model.User, error)
	InsertWallet(ctx context.Context, id int64) error
	CheckDataUser(ctx context.Context, username string) (*model.User, error)
	GetListPlayer(ctx context.Context) (*[]model.User, error)
	GetDetailPlayer(ctx context.Context, id int64) (*model.User, error)
	InsertAccount(ctx context.Context, data request2.CreateAccountRequest) error
	CheckIdBank(ctx context.Context, id int64) error
	CheckAccount(ctx context.Context, id int64) (*model.Account, error)
	UpdateBalance(ctx context.Context, id int64, total float64) error
	UpdateWallet(ctx context.Context, id int64, total float64) error
	GetWallet(ctx context.Context, id int64) (*model.Wallet, error)
	GetAccount(ctx context.Context, id int64) (*model.Account, error)
	GetBank(ctx context.Context, id int64) (*model.Bank, error)
	UpdateToken(ctx context.Context, id int64, token string) error
}
