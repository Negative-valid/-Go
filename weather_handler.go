package handlers

import (
	"fmt"
	"strings"
	"weather-bot/internal/api/openweathermap"
	"weather-bot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WeatherHandler struct {
	api       *openweathermap.Client
	logger    *logger.Logger
	cityCache map[int64]string
}

func NewWeatherHandler(api *openweathermap.Client, log *logger.Logger) *WeatherHandler {
	return &WeatherHandler{
		api:       api,
		logger:    log,
		cityCache: make(map[int64]string),
	}
}

func (h *WeatherHandler) Handle(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	if msg.IsCommand() {
		switch msg.Command() {
		case "start":
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Привіт! Напиши назву міста, щоб отримати прогноз погоди 🌤"))
		case "help":
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Інструкція:\n1. Напиши назву міста (наприклад: 'Київ').\n2. Якщо ти вже шукав місто, напиши 'погода', щоб отримати оновлені дані для останнього міста."))
		}
		return
	}

	city := strings.TrimSpace(msg.Text)
	if city == "" {
		return
	}

	if strings.ToLower(city) == "погода" {
		lastCity, exists := h.cityCache[msg.Chat.ID]
		if !exists {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ви ще не шукали погоду. Введіть назву міста."))
			return
		}
		city = lastCity
	} else {
		h.cityCache[msg.Chat.ID] = city
	}

	h.sendWeather(bot, msg.Chat.ID, city)
}

func (h *WeatherHandler) sendWeather(bot *tgbotapi.BotAPI, chatID int64, city string) {
	data, err := h.api.GetCurrentWeather(city)
	if err != nil {
		h.logger.Error("Помилка API: %v", err)
		bot.Send(tgbotapi.NewMessage(chatID, "❌ Не вдалося знайти погоду для цього міста. Перевірте назву."))
		return
	}

	text := fmt.Sprintf(
		"🌍 Погода в місті **%s**:\n"+
			"🌡 Температура: %.1f°C\n"+
			"💧 Вологість: %d%%\n"+
			"🌬 Вітер: %.1f м/с\n"+
			"📝 Опис: %s",
		data.Name, data.Main.Temp, data.Main.Humidity, data.Wind.Speed, data.Weather[0].Description,
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
