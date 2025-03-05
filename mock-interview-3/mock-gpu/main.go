package main

import (
	"fmt"
)

// Интерфейс GPU
type GPU interface {
	GetCurrentSpeed(model string) (CurrentSpeed, error)
	GetSpeedForecast(model string) ([]SpeedForecast, error)
}

// Структура для текущей скорости видеокарты
type CurrentSpeed struct {
	Model       string
	ClockSpeed  string // Тактовая частота
	MemorySpeed string // Скорость памяти
}

// Структура для прогноза скорости видеокарты
type SpeedForecast struct {
	Mode        string // Режим работы (например, "Игры", "Майнинг")
	ClockSpeed  string // Тактовая частота
	MemorySpeed string // Скорость памяти
}

// Мок-объект, реализующий интерфейс GPU
type MockGPU struct{}

// Метод для получения текущей скорости
func (m MockGPU) GetCurrentSpeed(model string) (CurrentSpeed, error) {
	return CurrentSpeed{
		Model:       model,
		ClockSpeed:  "1800 MHz",
		MemorySpeed: "14 Gbps",
	}, nil
}

// Метод для получения прогноза скорости
func (m MockGPU) GetSpeedForecast(model string) ([]SpeedForecast, error) {
	return []SpeedForecast{
		{
			Mode:        "Игры",
			ClockSpeed:  "1800 MHz",
			MemorySpeed: "14 Gbps",
		},
		{
			Mode:        "Майнинг",
			ClockSpeed:  "1200 MHz",
			MemorySpeed: "10 Gbps",
		},
		{
			Mode:        "Офис",
			ClockSpeed:  "800 MHz",
			MemorySpeed: "8 Gbps",
		},
	}, nil
}

func main() {
	// Создаем мок-объект
	mock := MockGPU{}

	// Получаем текущую скорость видеокарты
	currentSpeed, _ := mock.GetCurrentSpeed("RTX 3080")
	fmt.Println("Текущая скорость видеокарты:")
	fmt.Printf("Модель: %s\n", currentSpeed.Model)
	fmt.Printf("Тактовая частота: %s\n", currentSpeed.ClockSpeed)
	fmt.Printf("Скорость памяти: %s\n", currentSpeed.MemorySpeed)
	fmt.Println()

	// Получаем прогноз скорости
	forecast, _ := mock.GetSpeedForecast("RTX 3080")
	fmt.Println("Прогноз скорости:")
	for _, mode := range forecast {
		fmt.Printf("Режим: %s\n", mode.Mode)
		fmt.Printf("Тактовая частота: %s\n", mode.ClockSpeed)
		fmt.Printf("Скорость памяти: %s\n", mode.MemorySpeed)
		fmt.Println()
	}
}
