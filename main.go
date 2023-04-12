package main

import (
	"github.com/runan13/echo-rest-api/app"
	"github.com/runan13/echo-rest-api/app/routes"
	"github.com/runan13/echo-rest-api/config"
	"github.com/runan13/echo-rest-api/utils"
)

func main() {
	cfg := config.NewConfig()

	logger := utils.NewLogger()

	server := app.NewServer(cfg, logger)

	routes.Init(server, server.Logger)

	err := server.Start(cfg.HTTP.Port)
	if err != nil {
		server.Logger.Error().Msg("Port already used")
	}
}
