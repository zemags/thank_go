package main

import (
	"math/rand"
	"time"
)

// WeatherService предсказывает погоду.
type WeatherService interface {
	Forecast() int
}

// RealWeatherService представляет реальную службу WeatherService.
type RealWeatherService struct{}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *RealWeatherService) Forecast() int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(31)
	sign := rand.Intn(2)
	if sign == 1 {
		value = -value
	}
	return value
}

// Weather выдает текстовый прогноз погоды.
type Weather struct {
	service WeatherService
}

// Forecast сообщает текстовый прогноз погоды на завтра.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}

// Исправьте Weather, чтобы его можно было
// нормально протестировать. Сделайте заглушку
// WeatherService и используйте ее в TestForecast.

// Не меняйте WeatherService, метод Weather.Forecast()
// и набор тестов tests.
