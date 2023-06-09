package app

import (
	"apiserver/internal/middleware"
	"apiserver/pkg/config"
	"apiserver/pkg/server"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config
	Logger *zap.Logger
	Server *server.HttpServer
	Router server.Router
	DB     *gorm.DB
}

func NewApplication(config *config.Config, db *gorm.DB, logger *zap.Logger, server *server.HttpServer, router server.Router) *Application {
	return &Application{
		Config: config,
		Logger: logger,
		Server: server,
		Router: router,
		DB:     db,
	}
}

func (a *Application) Run() error {
	a.Server.Run(a.Router, middleware.NewMidddleware())

	return nil
}

func (a *Application) Close() {
	a.Logger.Info("perform application cleanup")
	sqlDB, err := a.DB.DB()
	if err != nil {
		a.Logger.Error("get sqlDB failed: %w", zap.Error(err))
	}
	err = sqlDB.Close()
	if err != nil {
		a.Logger.Error("get sqlDB failed: %w", zap.Error(err))
	}
}
