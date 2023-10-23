package repository

import (
	"context"
	"errors"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetListPlayer(ctx context.Context, userLevel string) (*[]model.User, error) {
	var users []model.User
	db := s.super.PostgresSql

	if userLevel != "" {
		db = db.Where("level = ?", userLevel)
	}

	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(users) == 0 {
		return nil, errors.New("user not found")
	}

	return &users, nil
}
