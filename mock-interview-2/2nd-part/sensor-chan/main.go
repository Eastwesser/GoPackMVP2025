package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CharacterData представляет структуру данных о персонаже.
type CharacterData struct {
	Character string  // Имя персонажа (например, "Diluc", "Venti", "Zhongli")
	Stat      string  // Характеристика (например, "HP", "Energy", "Attack")
	Value     float64 // Значение характеристики
}

// generateCharacterData генерирует случайные данные о персонаже и отправляет их в канал.
func generateCharacterData(character, stat string, ch chan<- CharacterData, delay time.Duration) {
	defer close(ch)                  // Закрытие канала после завершения генерации данных
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	for i := 0; i < 5; i++ {         // Генерация 5 значений
		time.Sleep(delay)                                                   // Задержка между генерацией значений
		value := rand.Float64() * 1000                                      // Генерация случайного значения от 0 до 1000
		ch <- CharacterData{Character: character, Stat: stat, Value: value} // Отправка данных в канал
	}
}

// processCharacterData обрабатывает данные о персонажах и возвращает канал с результатами.
func processCharacterData(ch ...chan CharacterData) chan string {
	var wg sync.WaitGroup    // Используется для ожидания завершения всех горутин
	out := make(chan string) // Создание канала для результатов обработки

	// Функция для обработки данных о персонаже
	process := func(c <-chan CharacterData) {
		defer wg.Done()       // Уменьшение счетчика WaitGroup при завершении горутины
		for data := range c { // Чтение данных из канала
			// Форматирование строки с результатом обработки
			out <- fmt.Sprintf("Обработаны данные: %s - %s: %.2f", data.Character, data.Stat, data.Value)
		}
	}

	wg.Add(len(ch)) // Увеличение счетчика WaitGroup на количество каналов
	for _, c := range ch {
		go process(c) // Запуск горутины для обработки данных от каждого канала
	}

	// Закрываем общий канал после завершения всех горутин
	go func() {
		wg.Wait()  // Ожидание завершения всех горутин
		close(out) // Закрытие канала с результатами
	}()

	return out
}

func main() {
	// Создание каналов для данных о персонажах
	dilucHPChan := make(chan CharacterData)
	ventiEnergyChan := make(chan CharacterData)
	zhongliAttackChan := make(chan CharacterData)

	// Запуск горутин для генерации данных о персонажах
	go generateCharacterData("Diluc", "HP", dilucHPChan, 100*time.Millisecond)
	go generateCharacterData("Venti", "Energy", ventiEnergyChan, 200*time.Millisecond)
	go generateCharacterData("Zhongli", "Attack", zhongliAttackChan, 300*time.Millisecond)

	// Обработка данных о персонажах
	processedDataChan := processCharacterData(dilucHPChan, ventiEnergyChan, zhongliAttackChan)

	// Чтение и вывод обработанных данных
	for data := range processedDataChan {
		fmt.Println(data)
	}

	fmt.Println("Программа завершена.")
}
