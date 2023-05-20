//go:build wireinject

package main

import (
	"apiserver/internal/app"
	"apiserver/internal/router"
	"apiserver/pkg/config"
	"apiserver/pkg/db"
	"apiserver/pkg/logger"
	"apiserver/pkg/server"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.New,
	logger.New,
	db.New,
	server.NewHttpServer,
	app.NewApplication,
	router.ProviderSet,
)

func createApp(configFile string) (*app.Application, func(), error) {
	panic(wire.Build(providerSet))
}
