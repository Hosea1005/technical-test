package main

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	dUsecase "service-auth/domain/usecase"
	"service-auth/internal/config"
	httpHandler "service-auth/internal/http"
	"service-auth/internal/repository/repository"
	"service-auth/internal/usecase"
)

func main() {
	logger := config.Logger()
	config.Environment()
	logger.Info("initializing service authorization ")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	repo := config.BaseRepository{
		PostgresSql: config.PostgresDB(logger),
		Redis:       config.Redis(ctx, logger),
	}
	// Initialize repository instances
	userRepo := repository.NewAuthRepository(repo, logger)
	cacheRepo := repository.NewCacheRepository(repo, logger)
	// Initialize use case instances
	userService := usecase.NewAuthUseCase(logger, userRepo, cacheRepo, repo)

	r := mux.NewRouter()
	initHandler(r, userService)
	http.Handle("/", r)

}
func initHandler(r *mux.Router, usecase dUsecase.AuthUseCase) {
	httpHandler.NewDeliveryHttpArea(r, usecase)
}
