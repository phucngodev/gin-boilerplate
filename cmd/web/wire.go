//go:build wireinject

package main

import (
	"apiserver/internal/app"
	"apiserver/internal/router"
	"apiserver/pkg/config"
	"apiserver/pkg/logger"
	"apiserver/pkg/server"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.New,
	logger.New,
	server.NewHttpServer,
	app.NewApplication,
	router.ProviderSet,
)

func createApp(configFile string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
