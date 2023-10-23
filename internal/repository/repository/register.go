package repository

import (
	"context"
	"service-auth/internal/http/register/request"
	"service-auth/internal/model"
)

func (s *AuthRepository) InsertUser(ctx context.Context, data request.RegisterRequest) (*model.User, error) {
	user := model.User{
		Username: data.Username,
		Password: data.Password,
		Fullname: data.Name,
		Level:    data.Level,
	}
	result := s.super.PostgresSql.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
