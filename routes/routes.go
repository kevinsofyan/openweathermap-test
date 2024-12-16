package routes

import (
	"weather/config"
	"weather/controllers"
	"weather/repositories"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	weatherRepo := repositories.NewWeatherRepository(config.DB)

	// Initialize controllers
	weatherController := controllers.NewWeatherController(weatherRepo)

	// Define routes
	e.GET("/weather", weatherController.FetchWeather)
	e.POST("/weather", weatherController.SaveWeather)
}
