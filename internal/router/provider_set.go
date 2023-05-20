package router

import (
	"apiserver/internal/handler"
	"apiserver/internal/repo/mysql"
	"apiserver/internal/service"
	"apiserver/pkg/server"

	"github.com/google/wire"
)

var userProviderSet = wire.NewSet(
	mysql.ProviderSet,
	service.ProviderSet,
	handler.NewUserHandler,
)

var ProviderSet = wire.NewSet(
	NewApiRouter,
	handler.NewHomeHandler,
	userProviderSet,
	wire.Bind(new(server.Router), new(*ApiRouter)),
)
