package main

import (
	"flag"
	"log"

	"go.uber.org/zap"
)

var configFile = flag.String("f", "config.yml", "set config file which viper will load")

func main() {
	flag.Parse()

	app, err := createApp(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	app.Server.RegisterOnShutdown(app.Close)
	if err := app.Run(); err != nil {
		app.Logger.Error("app.Run() errro: ", zap.Error(err))
	}
}
