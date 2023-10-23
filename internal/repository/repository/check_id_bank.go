package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) CheckIdBank(ctx context.Context, id int64) error {
	// Lakukan kueri ke database untuk memeriksa ID bank.
	// Misalnya, menggunakan GORM:

	var bank model.Bank // Gantilah 'Bank' dengan nama model yang sesuai
	result := s.super.PostgresSql.Where("id = ?", id).First(&bank)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("Bank not found")
		}
		return result.Error
	}

	// Jika ID bank ditemukan di database, return nil (tanpa kesalahan).
	return nil
}
