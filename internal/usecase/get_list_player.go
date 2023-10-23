package usecase

import (
	"context"
	lRequest "service-auth/internal/http/list_player/request"
	lResponse "service-auth/internal/http/list_player/response"
)

func (a AuthUseCase) GetListPlayer(ctx context.Context, request lRequest.ListPlayerRequest) *lResponse.ListPlayerResponse {

	_, errToken := a.CacheService.GetToken(ctx, request.Token)
	if errToken != nil {
		return &lResponse.ListPlayerResponse{
			Status: lResponse.Status{
				Code:    400,
				Message: "Unauthorized",
			},
		}
	}
	resUser, err := a.UserRepository.GetListPlayer(ctx, request.Search)
	if err != nil {
		return &lResponse.ListPlayerResponse{
			Status: lResponse.Status{
				Code:    400,
				Message: err.Error(),
			},
		}
	}
	var playerData []lResponse.Player
	for _, user := range *resUser {
		playerData = append(playerData, lResponse.Player{
			ID:        user.ID,
			Name:      user.Fullname,
			Username:  user.Username,
			Level:     user.Level,
			CreatedAt: user.CreatedAt,
		})
	}

	return &lResponse.ListPlayerResponse{
		Status: lResponse.Status{
			Code:    200,
			Message: "Success",
		},
		Data: lResponse.Data{
			Player: playerData,
		},
	}
}
