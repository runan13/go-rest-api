package services

import (
	"github.com/labstack/echo/v4"
	"github.com/runan13/echo-rest-api/types"
	"github.com/runan13/echo-rest-api/utils"
	"gorm.io/gorm"
	"net/http"
)

type RestaurantsService interface {
	GetAllRestaurants(c echo.Context) error
	PostRestaurant(c echo.Context) error
	DeleteRestaurant(c echo.Context) error
}

type restaurantsService struct {
	logger *utils.Logger
	db     *gorm.DB
}

func NewRestaurantsService(db *gorm.DB, logger *utils.Logger) RestaurantsService {
	return &restaurantsService{logger: logger, db: db}
}

func (s *restaurantsService) GetAllRestaurants(c echo.Context) error {
	res := &types.ErrorResponse{
		ErrorCode:    123,
		ErrorMessage: "123",
	}
	return c.JSON(http.StatusOK, res)
}

func (s *restaurantsService) PostRestaurant(c echo.Context) error {
	body := new(types.PostRestaurantRequest)
	_ = c.Bind(body)

	if err := c.Validate(body); err != nil {
		s.logger.Error().Err(err).Msg("Request POST /api/restaurants")
		return c.JSON(http.StatusBadRequest, &types.ErrorResponse{
			ErrorCode:    456,
			ErrorMessage: "Request body error",
		})
	}

	return c.JSON(http.StatusOK, body)
}

func (s *restaurantsService) DeleteRestaurant(c echo.Context) error {
	res := &types.ErrorResponse{
		ErrorCode:    123,
		ErrorMessage: "123",
	}
	return c.JSON(http.StatusOK, res)
}
