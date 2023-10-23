package repository

import (
	"context"
	"errors"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetListPlayer(ctx context.Context) (*[]model.User, error) {
	var users []model.User
	result := s.super.PostgresSql.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(users) == 0 {
		return nil, errors.New("No players found with the given name")
	}

	return &users, nil
}
