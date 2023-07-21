package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

type WeatherServiceStub struct {
	deg int
}

// Forecast возвращает фиксированное значение температуры (-5) для тестов.
func (stub *WeatherServiceStub) Forecast() int {
	return stub.deg
}

func TestForecast(t *testing.T) {
	stub := &WeatherServiceStub{}
	weather := Weather{service: stub}
	for _, test := range tests {
		stub.deg = test.deg
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
