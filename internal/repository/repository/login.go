package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) CheckDataUser(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	result := s.super.PostgresSql.Where("username = ?", username).First(&user)
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
