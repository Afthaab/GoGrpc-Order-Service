//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/order/service/pkg/api"
	"github.com/order/service/pkg/api/handler"
	"github.com/order/service/pkg/config"
	"github.com/order/service/pkg/db"
	"github.com/order/service/pkg/repository"
	"github.com/order/service/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	wire.Build(
		db.ConnectToDataBase,
		repository.NewOrderRepo,
		usecase.NewOrderUseCase,
		handler.NewOrderHandler,
		api.NewServerHttp)
	return &api.ServerHttp{}, nil
}

//go run github.com/google/wire/cmd/wire
