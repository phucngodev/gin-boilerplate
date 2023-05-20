package app

import (
	"apiserver/internal/middleware"
	"apiserver/pkg/config"
	"apiserver/pkg/server"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewApplication)

type Application struct {
	config *config.Config
	logger *zap.Logger
	server *server.HttpServer
	router server.Router
}

func NewApplication(config *config.Config, logger *zap.Logger, server *server.HttpServer, router server.Router) *Application {
	return &Application{
		config: config,
		logger: logger,
		server: server,
		router: router,
	}
}

func (a *Application) Run() error {
	a.server.Run(a.router, middleware.NewMidddleware())

	return nil
}