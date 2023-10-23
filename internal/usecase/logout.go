package usecase

import (
	"context"
	request4 "service-auth/internal/http/logout/request"
	response4 "service-auth/internal/http/logout/response"
)

func (a AuthUseCase) Logout(ctx context.Context, request request4.LogoutRequest) *response4.LogoutResponse {
	resUser, err := a.UserRepository.GetDetailPlayer(ctx, request.ID)
	if err != nil {
		return &response4.LogoutResponse{
			Status: response4.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	err = a.CacheService.DeleteToken(ctx, resUser.Token)
	if err != nil {
		return &response4.LogoutResponse{
			Status: response4.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	return &response4.LogoutResponse{
		Status: response4.Status{
			Code:    200,
			Message: "Success Login",
		},
	}
}
