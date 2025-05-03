package main

import (
	"context"
	"fmt"
	"sync"
)

// Worker маркирует ящики и отправляет их в канал
func worker(ctx context.Context, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик при завершении

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
			// Контекст отменён — завершаем работу
			fmt.Println("Worker: Контекст отменён, ухожу.")
			return
		}
	}
}

func main() {
	jobs := make(chan int)                                  // Канал для передачи чисел
	results := make(chan string)                            // Канал для передачи маркированных значений
	ctx, cancel := context.WithCancel(context.Background()) // Создаём контекст с возможностью отмены
	defer cancel()                                          // Гарантируем отмену контекста при завершении

	var wg sync.WaitGroup

	// Запускаем worker в отдельной горутине
	wg.Add(1)
	go worker(ctx, jobs, results, &wg)

	// Наполняем канал задачами
	go func() {
		for i := 1; i <= 4; i++ {
			jobs <- i
			fmt.Println("Отправлена задача:", i)
		}
		close(jobs) // Закрываем канал после отправки всех задач
	}()

	// Читаем результаты, пока results не закроется
	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range results {
			fmt.Println("Результат:", res)
		}
	}()

	// Ждём завершения всех горутин
	wg.Wait()
	fmt.Println("Все задачи завершены.")
}
