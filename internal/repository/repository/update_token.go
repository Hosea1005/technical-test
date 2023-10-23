package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) UpdateToken(ctx context.Context, id int64, token string) error {
	var data model.User
	result := s.super.PostgresSql.First(&data, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("Data not found")
		}
		return result.Error
	}

	// Perbarui token
	data.Token = token
	result = s.super.PostgresSql.Save(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
