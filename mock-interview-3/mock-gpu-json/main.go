package main

import (
	"encoding/json"
	"fmt"
)

// Интерфейс GPU
type GPU interface {
	GetCurrentSpeed(model string) (CurrentSpeed, error)
	GetSpeedForecast(model string) ([]SpeedForecast, error)
}

// Структура для текущей скорости видеокарты
type CurrentSpeed struct {
	Model       string `json:"model"`
	ClockSpeed  string `json:"clock_speed"`
	MemorySpeed string `json:"memory_speed"`
}

// Структура для прогноза скорости видеокарты
type SpeedForecast struct {
	Mode        string `json:"mode"`
	ClockSpeed  string `json:"clock_speed"`
	MemorySpeed string `json:"memory_speed"`
}

// Мок-объект, реализующий интерфейс GPU
type MockGPU struct{}

// Метод для получения текущей скорости
func (m MockGPU) GetCurrentSpeed(model string) (CurrentSpeed, error) {
	// JSON-строка с данными о текущей скорости
	jsonCurrentSpeed := `{"model":"RTX 3080","clock_speed":"1800 MHz","memory_speed":"14 Gbps"}`

	// Десериализация JSON в структуру CurrentSpeed
	var currentSpeed CurrentSpeed
	err := json.Unmarshal([]byte(jsonCurrentSpeed), &currentSpeed)
	if err != nil {
		return CurrentSpeed{}, err
	}

	return currentSpeed, nil
}

// Метод для получения прогноза скорости
func (m MockGPU) GetSpeedForecast(model string) ([]SpeedForecast, error) {
	// JSON-строка с данными о прогнозе скорости
	jsonForecast := `[
		{"mode":"Игры","clock_speed":"1800 MHz","memory_speed":"14 Gbps"},
		{"mode":"Майнинг","clock_speed":"1200 MHz","memory_speed":"10 Gbps"},
		{"mode":"Офис","clock_speed":"800 MHz","memory_speed":"8 Gbps"}
	]`

	// Десериализация JSON в срез структур SpeedForecast
	var forecast []SpeedForecast
	err := json.Unmarshal([]byte(jsonForecast), &forecast)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

func main() {
	// Создаем мок-объект
	mock := MockGPU{}

	// Получаем текущую скорость видеокарты
	currentSpeed, err := mock.GetCurrentSpeed("RTX 3080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Текущая скорость видеокарты:")
	fmt.Printf("Модель: %s\n", currentSpeed.Model)
	fmt.Printf("Тактовая частота: %s\n", currentSpeed.ClockSpeed)
	fmt.Printf("Скорость памяти: %s\n", currentSpeed.MemorySpeed)
	fmt.Println()

	// Получаем прогноз скорости
	forecast, err := mock.GetSpeedForecast("RTX 3080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Прогноз скорости:")
	for _, mode := range forecast {
		fmt.Printf("Режим: %s\n", mode.Mode)
		fmt.Printf("Тактовая частота: %s\n", mode.ClockSpeed)
		fmt.Printf("Скорость памяти: %s\n", mode.MemorySpeed)
		fmt.Println()
	}
}
