package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"service-auth/internal/model"
)

func (s *AuthRepository) UpdateWallet(ctx context.Context, id int64, total float64) error {
	// Gunakan GORM untuk menemukan akun berdasarkan ID
	var wallet model.Wallet
	result := s.super.PostgresSql.First(&wallet, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("account not found")
		}
		return result.Error
	}

	// Perbarui saldo akun
	wallet.Balance = total
	result = s.super.PostgresSql.Save(&wallet)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
