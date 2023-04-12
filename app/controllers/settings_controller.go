package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/runan13/echo-rest-api/app/services"
	"github.com/runan13/echo-rest-api/types"
	"net/http"
)

type SettingsController interface {
	GetLocales(c echo.Context) error
	PostLocales(c echo.Context) error
}

type settingsController struct {
	service services.SettingsService
}

func NewSettingsController() *settingsController {
	return &settingsController{service: services.NewSettingsService()}
}

func (controller *settingsController) GetLocales(c echo.Context) error {
	res := &types.ErrorResponse{
		ErrorCode:    123,
		ErrorMessage: "123",
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *settingsController) PostLocales(c echo.Context) error {
	res := &types.ErrorResponse{
		ErrorCode:    123,
		ErrorMessage: "123",
	}
	return c.JSON(http.StatusOK, res)
}
