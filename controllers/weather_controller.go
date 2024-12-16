package controllers

import (
	"net/http"
	"weather/models"
	"weather/repositories"

	"github.com/labstack/echo/v4"
)

type WeatherController struct {
	repo repositories.WeatherRepository
}

func NewWeatherController(repo repositories.WeatherRepository) *WeatherController {
	return &WeatherController{repo}
}

// @Summary Fetch weather data
// @Description Fetch weather data for a given city from the OpenWeatherMap API
// @Tags weather
// @Accept json
// @Produce json
// @Param city query string true "City"
// @Success 200 {object} models.Weather
// @Failure 500 {object} ErrorResponse
// @Router /weather [get]
func (ctrl *WeatherController) FetchWeather(c echo.Context) error {
	city := c.QueryParam("city")
	if city == "" {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "City is required"})
	}

	weather, err := ctrl.repo.FetchWeather(city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to fetch weather data"})
	}

	return c.JSON(http.StatusOK, weather)
}

// @Summary Save weather data
// @Description Save weather data to the database
// @Tags weather
// @Accept json
// @Produce json
// @Param weather body models.Weather true "Weather"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} ErrorResponse
// @Router /weather [post]
func (ctrl *WeatherController) SaveWeather(c echo.Context) error {
	weather := new(models.Weather)
	if err := c.Bind(weather); err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid request"})
	}

	if err := ctrl.repo.SaveWeather(weather); err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Failed to save weather data"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Weather data saved successfully",
	})
}
