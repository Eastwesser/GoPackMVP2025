package main

import (
	"context"
	"fmt"
	"time"
)

// Worker маркирует ящики и отправляет их в канал
func worker(ctx context.Context, jobs <-chan int, results chan<- string) {
	for {
		select {
		case num, ok := <-jobs:
			if !ok {
				// Канал закрыт — завершаем работу
				fmt.Println("Worker: Канал закрыт, ухожу.")
				close(results)
				return
			}
			fmt.Println("Worker: обработал", num)
			results <- fmt.Sprintf("Eastwesser-%d", num)

		case <-ctx.Done():

			fmt.Println("Worker: Контекст отменён, ухожу.") // Контекст отменён — завершаем работу
			return
		}
	}
}

func main() {
	jobs := make(chan int)                                  // канал для передачи чисел
	results := make(chan string)                            // канал для передачи маркированных значений
	ctx, cancel := context.WithCancel(context.Background()) //создаём контекст с возможностью отмены (cancel()).

	go worker(ctx, jobs, results) // Запускаем worker в отдельной горутине. Теперь worker ждёт данные из jobs.

	// Наполняем канал задачами
	go func() {
		for i := 1; i <= 4; i++ {
			jobs <- i
			fmt.Println(i)
		}
		// Закрываем канал jobs (но Worker ещё ждёт!)
		close(jobs) // Закрываем канал сразу после завершения цикла!
	}()

	// Симуляция зависания: Worker остаётся ждать
	//time.Sleep(1000 * time.Millisecond) // "Я якобы поработал"

	/*
		Ждём 1 секунду (1000 мс).
		Это имитация "зависшей" горутины.
		worker ждёт новые задания в jobs, но новых чисел уже нет.
		Он не завершится сам, так как канал jobs не закрыт.
	*/

	// Читаем результаты, пока results не закроется
	go func() {
		for res := range results {
			fmt.Println("Результат:", res)
		}
	}()

	//close(jobs) // Опа! Надо закрывать канал!

	// Дадим Worker обработать все задачи
	time.Sleep(2 * time.Second) // Ждём 2 секунды, чтобы worker успел обработать оставшиеся данные.

	// Завершаем контекст (если бы канал не закрылся)
	cancel()
}
