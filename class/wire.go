//go:build wireinject
// +build wireinject

package main

import (
	"20241112/handler"
	"20241112/repository"
	"20241112/service"
	"github.com/google/wire"
)

func InitializeOrderHandler() handler.OrderHandler {
	wire.Build(handler.InitOrderHandler, service.InitOrderService, repository.InitOrderRepo)
	return handler.OrderHandler{}
}
