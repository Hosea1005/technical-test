package usecase

import (
	"go.uber.org/zap"
	domain "service-auth/domain/repository"
	"service-auth/domain/usecase"
	"service-auth/internal/config"
)

type AuthUseCase struct {
	Logger         *zap.Logger
	UserRepository domain.AuthRepository
	CacheService   domain.CacheRepository
	Super          config.BaseRepository
}

func NewAuthUseCase(
	logger *zap.Logger,
	userRepository domain.AuthRepository,
	cacheService domain.CacheRepository,
	super config.BaseRepository,
) usecase.AuthUseCase {
	return &AuthUseCase{Logger: logger, UserRepository: userRepository, CacheService: cacheService, Super: super}
}
