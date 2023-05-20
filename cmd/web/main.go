package main

import (
	"flag"
	"log"
)

var configFile = flag.String("f", "config.yml", "set config file which viper will load")

func main() {
	flag.Parse()

	app, err := createApp(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
