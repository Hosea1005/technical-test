package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetWallet(ctx context.Context, id int64) (*model.Wallet, error) {
	var wallet model.Wallet
	result := s.super.PostgresSql.Where("id_user = ?", id).First(&wallet)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &model.Wallet{}, nil
		}
		return nil, result.Error
	}
	return &wallet, nil
}
