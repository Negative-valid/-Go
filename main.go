package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

const APIKey = "TOKEN"

func getCurrentWeather(city string) (*WeatherResponse, error) {
	safeCity := url.QueryEscape(city)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ua",
		safeCity, APIKey)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("помилка мережі: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API повернув помилку: %d (перевірте назву міста)", resp.StatusCode)
	}

	var w WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return nil, fmt.Errorf("помилка обробки JSON: %v", err)
	}

	return &w, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введіть назву міста (наприклад, Вінниця): ")

	input, _ := reader.ReadString('\n')
	city := strings.TrimSpace(input)

	if city == "" {
		fmt.Println("Ви не ввели назву міста.")
		return
	}

	fmt.Println("Отримання даних...")
	data, err := getCurrentWeather(city)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	desc := "невідомо"
	if len(data.Weather) > 0 {
		desc = data.Weather[0].Description
	}

	fmt.Printf("🌍 Погода в місті %s:\n", data.Name)
	fmt.Printf("🌡  Температура:    %.1f°C\n", data.Main.Temp)
	fmt.Printf("🤔 Відчувається як: %.1f°C\n", data.Main.FeelsLike)
	fmt.Printf("💧 Вологість:      %d%%\n", data.Main.Humidity)
	fmt.Printf("🌬  Вітер:          %.1f м/с\n", data.Wind.Speed)
	fmt.Printf("📝 Опис:           %s\n", desc)
}
