package main

import (
	"fmt"
	"time"
)

// Функция tryReceive пытается получить значение из канала без блокировки.
func tryReceive(ch chan int) (value int, ok bool) {
	select {
	case value, ok = <-ch: // Попытка получить значение из канала
		if !ok {
			fmt.Println("Канал закрыт")
			return 0, false
		}
		return value, true
	default:
		// Если канал пуст, возвращаем false
		return 0, false
	}
}

func main() {
	channelOne := make(chan int, 1) // Буферизованный канал с емкостью 1

	// Горутина, которая отправляет значение в канал через 2 секунды
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- 42 // Отправляем значение в канал
	}()

	// Пытаемся получить значение из канала несколько раз
	for i := 0; i < 5; i++ {
		value, ok := tryReceive(channelOne)
		if ok {
			fmt.Printf("Значение успешно получено из канала: %d\n", value)
		} else {
			fmt.Println("Значение не получено из канала. Он пуст или закрыт.")
		}
		time.Sleep(1 * time.Second) // Ждем 1 секунду перед следующей попыткой
	}

	// Закрываем канал после завершения
	close(channelOne)

	// Пытаемся получить значение из закрытого канала
	value, ok := tryReceive(channelOne)
	if ok {
		fmt.Printf("Значение успешно получено из канала: %d\n", value)
	} else {
		fmt.Println("Значение не получено из канала. Он пуст или закрыт.")
	}
}
