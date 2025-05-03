package main

import (
	"fmt"
	"sync"
	"time"
)

// generateActions генерирует действия персонажей и отправляет их в канал.
func generateGenshinActions(characterName string, actions []string, ch chan<- string) {
	defer close(ch) // Закрываем канал после завершения генерации действий
	for _, action := range actions {
		ch <- fmt.Sprintf("%s: %s", characterName, action) // Отправляем действие в канал
		time.Sleep(time.Millisecond * 200)                 // Имитируем задержку
	}
}

// aggregateActions объединяет данные из нескольких источников (каналов) в один канал.
func aggregateGenshinActions(done chan struct{}, sources ...chan string) chan string {
	var wg sync.WaitGroup    // Используем WaitGroup для синхронизации горутин
	out := make(chan string) // Создаем общий канал для объединенных данных

	// Функция для чтения из одного источника и записи в общий канал
	output := func(c <-chan string) {
		defer wg.Done()      // Уменьшаем счетчик WaitGroup при завершении горутины
		for msg := range c { // Читаем сообщения из канала
			out <- msg // Записываем сообщения в общий канал
		}
	}

	wg.Add(len(sources)) // Увеличиваем счетчик WaitGroup на количество источников
	for _, c := range sources {
		go output(c) // Запускаем горутину для каждого источника
	}

	// Закрываем общий канал после завершения всех горутин
	go func() {
		wg.Wait()   // Ждем завершения всех горутин
		close(out)  // Закрываем общий канал
		close(done) // Сигнализируем о завершении работы
	}()

	return out
}

func main() {
	// Создаем каналы для действий персонажей
	dilucActions := make(chan string)
	ventiActions := make(chan string)
	zhongliActions := make(chan string)

	// Канал для сигнала о завершении работы
	done := make(chan struct{})

	// Объединяем данные из всех источников в один канал
	aggregated := aggregateGenshinActions(done, dilucActions, ventiActions, zhongliActions)

	// Запускаем горутины для генерации действий персонажей
	go generateGenshinActions("Diluc", []string{"Атака мечом", "Огненный шторм", "Феникс"}, dilucActions)
	go generateGenshinActions("Venti", []string{"Выстрел из лука", "Небесная песнь", "Вихрь"}, ventiActions)
	go generateGenshinActions("Zhongli", []string{"Каменный щит", "Метеорит", "Падение кометы"}, zhongliActions)

	// Читаем и выводим действия из объединенного канала
	go func() {
		for msg := range aggregated {
			fmt.Println(msg)
		}
	}()

	// Ждем сигнала о завершении работы
	<-done
	fmt.Println("Все действия персонажей обработаны.")
}
