package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	WeatherToken  string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	tgToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if tgToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN не знайдено")
	}

	weatherToken := os.Getenv("OPENWEATHER_API_KEY")
	if weatherToken == "" {
		return nil, fmt.Errorf("OPENWEATHER_API_KEY не знайдено")
	}

	return &Config{
		TelegramToken: tgToken,
		WeatherToken:  weatherToken,
	}, nil
}