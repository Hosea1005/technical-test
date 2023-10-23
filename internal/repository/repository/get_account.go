package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetAccount(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account
	result := s.super.PostgresSql.Where("id_user = ?", id).First(&account)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &model.Account{}, nil
		}
		return nil, result.Error
	}
	return &account, nil
}
