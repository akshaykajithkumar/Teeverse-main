//go:build wireinject
// +build wireinject

package di

import (
	http "Teeverse/pkg/api"
	handler "Teeverse/pkg/api/handler"
	config "Teeverse/pkg/config"
	db "Teeverse/pkg/db"
	repository "Teeverse/pkg/repository"
	usecase "Teeverse/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, http.NewServerHTTP, repository.NewCategoryRepository, usecase.NewCategoryUseCase, handler.NewCategoryHandler, repository.NewInventoryRepository, usecase.NewInventoryUseCase, handler.NewInventoryHandler, repository.NewWishlistRepository, usecase.NewWishlistUseCase, handler.NewWishlistHandler, repository.NewOtpRepository, usecase.NewOtpUseCase, handler.NewOtpHandler, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, repository.NewCartRepository, usecase.NewCartUseCase, handler.NewCartHandler, repository.NewOrderRepository, usecase.NewOrderUseCase, handler.NewOrderHandler, repository.NewAdminRepository, usecase.NewAdminUseCase, handler.NewAdminHandler, repository.NewPaymentRepository, usecase.NewPaymentUseCase, handler.NewPaymentHandler)

	return &http.ServerHTTP{}, nil
}
