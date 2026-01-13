package main

import (
	"log"
	"weather-bot/internal/api/openweathermap"
	"weather-bot/internal/bot/handlers"
	"weather-bot/pkg/config"
	"weather-bot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	appLogger := logger.New()

	weatherClient := openweathermap.NewClient(cfg.WeatherToken)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		appLogger.Error("Bot init error: %v", err)
		log.Fatal(err)
	}
	bot.Debug = true

	appLogger.Info("Authorized as %s", bot.Self.UserName)

	weatherHandler := handlers.NewWeatherHandler(weatherClient, appLogger)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			weatherHandler.Handle(bot, update.Message)
		}
	}
}
