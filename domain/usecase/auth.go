package usecase

import (
	"context"
	request2 "service-auth/internal/http/create_account/request"
	response2 "service-auth/internal/http/create_account/response"
	dRequest "service-auth/internal/http/detail_player/request"
	dResponse "service-auth/internal/http/detail_player/response"
	lRequest "service-auth/internal/http/list_player/request"
	lResponse "service-auth/internal/http/list_player/response"
	"service-auth/internal/http/login/request"
	"service-auth/internal/http/login/response"
	request4 "service-auth/internal/http/logout/request"
	response4 "service-auth/internal/http/logout/response"
	rRequest "service-auth/internal/http/register/request"
	rResponse "service-auth/internal/http/register/response"
	request3 "service-auth/internal/http/top_up/request"
	response3 "service-auth/internal/http/top_up/response"
)

type AuthUseCase interface {
	Login(ctx context.Context, request request.LoginRequest) *response.LoginResponse
	Register(ctx context.Context, request rRequest.RegisterRequest) *rResponse.RegisterResponse
	Logout(ctx context.Context, request request4.LogoutRequest) *response4.LogoutResponse
	CreateAccount(ctx context.Context, request request2.CreateAccountRequest) *response2.CreateAccountResponse
	GetListPlayer(ctx context.Context, request lRequest.ListPlayerRequest) *lResponse.ListPlayerResponse
	GetDetailPlayer(ctx context.Context, request dRequest.DetailPlayerRequest) *dResponse.DetailPlayerResponse
	TopUp(ctx context.Context, request request3.TopUpRequest) *response3.TopUpResponse
}
