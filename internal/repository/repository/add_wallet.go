package repository

import (
	"context"
	"service-auth/internal/model"
)

func (s *AuthRepository) InsertWallet(ctx context.Context, id int64) error {
	user := model.Wallet{
		IDUser:  id,
		Balance: 0,
	}
	result := s.super.PostgresSql.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
