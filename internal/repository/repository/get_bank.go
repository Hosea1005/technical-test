package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetBank(ctx context.Context, id int64) (*model.Bank, error) {
	var bank model.Bank
	result := s.super.PostgresSql.Where("id = ?", id).First(&bank)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &model.Bank{}, nil
		}
		return nil, result.Error
	}
	return &bank, nil
}
