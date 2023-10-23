package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) GetDetailPlayer(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	result := s.super.PostgresSql.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, result.Error
	}

	// Di sini, Anda dapat memeriksa kata sandi sesuai dengan kebutuhan aplikasi Anda.
	// Misalnya, Anda bisa membandingkan password yang diberikan dengan password yang tersimpan di basis data.

	return &user, nil
}
