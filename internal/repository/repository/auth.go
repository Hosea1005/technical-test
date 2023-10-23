package repository

import (
	"go.uber.org/zap"
	"service-auth/internal/config"
)

type AuthRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}

func NewAuthRepository(super config.BaseRepository,
	logger *zap.Logger) *AuthRepository {
	return &AuthRepository{
		super:  super,
		logger: logger,
	}
}
func NewCacheRepository(super config.BaseRepository,
	logger *zap.Logger) *CacheRepository {
	return &CacheRepository{
		super:  super,
		logger: logger,
	}
}

type CacheRepository struct {
	super  config.BaseRepository
	logger *zap.Logger
}
