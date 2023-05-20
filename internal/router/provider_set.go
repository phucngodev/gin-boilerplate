package router

import (
	"apiserver/internal/handler"
	"apiserver/pkg/server"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewApiRouter,
	handler.NewHomeHandler,
	wire.Bind(new(server.Router), new(*ApiRouter)),
)
