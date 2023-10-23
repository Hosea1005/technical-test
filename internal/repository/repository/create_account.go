package repository

import (
	"context"
	request2 "service-auth/internal/http/create_account/request"
	"service-auth/internal/model"
)

func (s *AuthRepository) InsertAccount(ctx context.Context, data request2.CreateAccountRequest) error {
	account := model.Account{
		Grade:   data.Grade,
		IDBank:  data.IDBank,
		IDUser:  data.IDUser,
		Balance: data.Balance,
	}
	result := s.super.PostgresSql.Create(&account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
