package main

import (
	"fmt"
	"sync"
)

// Artifact представляет информацию об артефакте.
type Artifact struct {
	Type   string // Тип артефакта (например, "Цветок жизни", "Перо смерти")
	Rarity int    // Редкость артефакта (например, 3, 4, 5)
}

// readFromChan читает данные из канала и возвращает слайс с этими данными.
func readFromChan(ch <-chan Artifact, count int) []Artifact {
	var result []Artifact
	for i := 0; i < count; i++ {
		value := <-ch                  // Чтение данных из канала
		result = append(result, value) // Добавление данных в слайс
	}
	return result
}

// mergeChannels объединяет 4 канала в один и выводит значения.
func mergeChannels(ch1, ch2, ch3, ch4 <-chan Artifact) <-chan Artifact {
	out := make(chan Artifact) // Создаем общий канал для объединенных данных
	var wg sync.WaitGroup      // Используем WaitGroup для синхронизации горутин

	// Функция для чтения из одного канала и записи в общий канал
	merge := func(c <-chan Artifact) {
		defer wg.Done()        // Уменьшаем счетчик WaitGroup при завершении горутины
		for value := range c { // Читаем данные из канала
			out <- value // Записываем данные в общий канал
		}
	}

	wg.Add(4)     // Увеличиваем счетчик WaitGroup на количество каналов
	go merge(ch1) // Запускаем горутину для каждого канала
	go merge(ch2)
	go merge(ch3)
	go merge(ch4)

	// Закрываем общий канал после завершения всех горутин
	go func() {
		wg.Wait()  // Ждем завершения всех горутин
		close(out) // Закрываем общий канал
	}()

	return out
}

func main() {
	// Создаем 4 канала для артефактов
	ch1 := make(chan Artifact)
	ch2 := make(chan Artifact)
	ch3 := make(chan Artifact)
	ch4 := make(chan Artifact)

	// Запускаем горутины для записи данных в каналы
	go func() {
		defer close(ch1)
		ch1 <- Artifact{Type: "Цветок жизни", Rarity: 5}
		ch1 <- Artifact{Type: "Перо смерти", Rarity: 4}
	}()
	go func() {
		defer close(ch2)
		ch2 <- Artifact{Type: "Пески времени", Rarity: 3}
		ch2 <- Artifact{Type: "Кубок пространства", Rarity: 5}
	}()
	go func() {
		defer close(ch3)
		ch3 <- Artifact{Type: "Корона разума", Rarity: 4}
		ch3 <- Artifact{Type: "Цветок жизни", Rarity: 3}
	}()
	go func() {
		defer close(ch4)
		ch4 <- Artifact{Type: "Перо смерти", Rarity: 5}
		ch4 <- Artifact{Type: "Пески времени", Rarity: 4}
	}()

	// Читаем данные из первого канала и выводим слайс
	data := readFromChan(ch1, 2)
	fmt.Println("Данные из ch1:")
	for _, artifact := range data {
		fmt.Printf("Тип: %s, Редкость: %d\n", artifact.Type, artifact.Rarity)
	}

	// Объединяем 4 канала в один
	mergedChan := mergeChannels(ch1, ch2, ch3, ch4)

	// Читаем и выводим данные из объединенного канала
	fmt.Println("Объединенные данные:")
	for artifact := range mergedChan {
		fmt.Printf("Тип: %s, Редкость: %d\n", artifact.Type, artifact.Rarity)
	}

	fmt.Println("Программа завершена.")
}
