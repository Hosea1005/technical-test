package domain

import (
	"context"
	"service-auth/internal/model"
)

type CacheRepository interface {
	Set(ctx context.Context, key string, value *model.TopUp) error
	SetToken(ctx context.Context, key string, token string) error
	GetToken(ctx context.Context, key string) (string, error)
	DeleteToken(ctx context.Context, key string) error
}
