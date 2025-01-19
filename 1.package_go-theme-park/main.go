package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ParkName    = "Happy Fun Park"
	MaxVisitors = 1000
	TicketPrice = 20.5
)

var (
	OpenTime  = time.Date(2025, time.January, 19, 9, 0, 0, 0, time.Local)
	CloseTime = time.Date(2025, time.January, 19, 22, 0, 0, 0, time.Local)
)

func main() {
	// Приветствие и базовая информация
	fmt.Println("Добро пожаловать в", ParkName)
	fmt.Printf("Часы работы: %v - %v\n", OpenTime.Format("15:04"), CloseTime.Format("15:04"))

	// Текущие данные
	visitorCount := 500
	fmt.Println("Текущие посетители:", visitorCount)

	// Выручка
	totalRevenue := calculateRevenue(visitorCount, TicketPrice)
	fmt.Printf("Общая выручка: %.2f$\n", totalRevenue)

	// Аттракционы
	attractions := []string{"Колесо обозрения", "Американские горки", "Комната страха", "Автодром"}
	printAttractions(attractions)

	// Статистика посетителей
	visitorAgeGroups := map[string]int{"Дети": 200, "Взрослые": 250, "Пенсионеры": 50}
	printVisitorStats(visitorAgeGroups)

	// Работа с указателем
	highlightAttraction := "Американские горки"
	promoteAttraction(&highlightAttraction)

	// Преобразование в строку
	visitorCountStr := strconv.Itoa(visitorCount)
	fmt.Println("Посетителей (строка):", visitorCountStr)

	// Описание парка
	parkDescription := strings.Join(attractions, ", ")
	fmt.Println("Описание парка:", parkDescription)

	// Проверка вместимости
	err := checkCapacity(visitorCount + 600)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Специальное объявление
	func(message string) {
		fmt.Println("Специальное объявление:", message)
	}("Скидки на билеты после 18:00!")

	// Логирование
	logToFile("park.log", "Парк успешно открылся!")

	fmt.Println("\n*** Парк завершил работу! ***")
	os.Exit(0)
}

// calculateRevenue рассчитывает выручку.
func calculateRevenue(visitors int, price float64) float64 {
	return float64(visitors) * price
}

// printAttractions выводит список аттракционов.
func printAttractions(attractions []string) {
	fmt.Println("Доступные аттракционы:")
	for _, attraction := range attractions {
		fmt.Println("- ", attraction)
	}
}

// printVisitorStats выводит статистику посетителей.
func printVisitorStats(stats map[string]int) {
	fmt.Println("Статистика посетителей:")
	for group, count := range stats {
		fmt.Printf("%s: %d\n", group, count)
	}
}

// promoteAttraction улучшает название аттракциона.
func promoteAttraction(attraction *string) {
	*attraction += " - САМЫЙ ПОПУЛЯРНЫЙ!"
	fmt.Println("Промоция:", *attraction)
}

// checkCapacity проверяет вместимость парка.
func checkCapacity(currentVisitors int) error {
	if currentVisitors > MaxVisitors {
		return errors.New("превышена максимальная вместимость парка")
	}
	return nil
}

// logToFile записывает сообщение в файл.
func logToFile(filename, message string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка записи лога:", err)
		return
	}
	defer f.Close()

	buffer := bytes.NewBufferString(message)
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		fmt.Println("Ошибка записи лога:", err)
	}
}
