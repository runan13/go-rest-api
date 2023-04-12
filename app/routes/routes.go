package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/runan13/echo-rest-api/app"
	"github.com/runan13/echo-rest-api/app/controllers"
	"github.com/runan13/echo-rest-api/utils"
)

func Init(s *app.Server, logger *utils.Logger) {
	setSettingsController(s, logger)
	setRestaurantsController(s, logger)
}

func setSettingsController(s *app.Server, logger *utils.Logger) {
	settings := controllers.NewSettingsController()
	s.Echo.GET("/api/settings/locale", func(c echo.Context) error { return settings.GetLocales(c) })
	s.Echo.POST("/api/settings/locale", func(c echo.Context) error { return settings.PostLocales(c) })
}

func setRestaurantsController(s *app.Server, logger *utils.Logger) {
	restaurants := controllers.NewRestaurantsController(s.DB, logger)
	s.Echo.GET("/api/restaurants", func(c echo.Context) error { return restaurants.GetAllRestaurants(c) })
	s.Echo.POST("/api/restaurants", func(c echo.Context) error { return restaurants.PostRestaurant(c) })
	s.Echo.DELETE("/api/restaurants", func(c echo.Context) error { return restaurants.DeleteRestaurant(c) })
}
