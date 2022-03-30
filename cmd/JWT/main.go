package main

import (
	"authentication/config"
	"authentication/internal/server"
	"log"
)

func main() {
	config.NewConfig()

	app := server.NewServer()

	if err := app.StartServer(config.Conf.Server.ListenAddress); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
