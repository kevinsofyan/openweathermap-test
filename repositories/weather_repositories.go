package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"weather/models"

	"gorm.io/gorm"
)

type WeatherRepository interface {
	FetchWeather(city string) (*models.Weather, error)
	SaveWeather(weather *models.Weather) error
}

type weatherRepository struct {
	db *gorm.DB
}

func NewWeatherRepository(db *gorm.DB) WeatherRepository {
	return &weatherRepository{db}
}

func (r *weatherRepository) FetchWeather(city string) (*models.Weather, error) {
	apiKey := os.Getenv("5e8ef125d69d1598fa0f692668f25dcc")
	if apiKey == "" {
		return nil, errors.New("API key is missing")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch weather data")
	}

	var weather models.Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func (r *weatherRepository) SaveWeather(weather *models.Weather) error {
	return r.db.Create(weather).Error
}
