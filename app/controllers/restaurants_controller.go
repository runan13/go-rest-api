package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/runan13/echo-rest-api/app/services"
	"github.com/runan13/echo-rest-api/utils"
	"gorm.io/gorm"
)

type RestaurantsController interface {
	GetAllRestaurants(c echo.Context) error
	PostRestaurant(c echo.Context) error
	DeleteRestaurant(c echo.Context) error
}

type restaurantsController struct {
	service services.RestaurantsService
}

func NewRestaurantsController(db *gorm.DB, logger *utils.Logger) *restaurantsController {
	return &restaurantsController{service: services.NewRestaurantsService(db, logger)}
}

func (controller *restaurantsController) GetAllRestaurants(c echo.Context) error {
	return controller.service.GetAllRestaurants(c)
}

func (controller *restaurantsController) PostRestaurant(c echo.Context) error {
	return controller.service.PostRestaurant(c)
}

func (controller *restaurantsController) DeleteRestaurant(c echo.Context) error {
	return controller.service.DeleteRestaurant(c)
}
