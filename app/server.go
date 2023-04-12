package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/runan13/echo-rest-api/app/database"
	"github.com/runan13/echo-rest-api/config"
	"github.com/runan13/echo-rest-api/utils"
	"gorm.io/gorm"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
	Logger *utils.Logger
}

func NewServer(cfg *config.Config, logger *utils.Logger) *Server {
	return &Server{
		Echo:   echo.New(),
		DB:     database.Init(cfg),
		Config: cfg,
		Logger: logger,
	}
}

func (server *Server) Start(addr string) error {
	server.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			switch v.Status {
			case 200:
				server.Logger.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Msg("request")
			case 400:
				server.Logger.Error().
					Str("URI", v.URI).
					Int("status", v.Status).
					Err(v.Error).
					Msg("request")
			}

			return nil
		},
	}))

	server.Echo.Validator = &utils.CustomValidator{Validator: validator.New()}

	return server.Echo.Start(":" + addr)
}
