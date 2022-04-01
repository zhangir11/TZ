package main

import (
	"authentication/config"
	server "authentication/internal/serverv2"
	"log"
)

func main() {
	config.NewConfig()

	app := server.NewServer()

	if err := app.StartServer(config.Conf.Server.ListenAddress); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
