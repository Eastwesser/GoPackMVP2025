package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Worker обрабатывает задачи из канала jobs и отправляет результаты в канал results
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик WaitGroup при завершении работы воркера

	for job := range jobs {
		fmt.Printf("Воркер %d начал обработку задачи %d\n", id, job)
		time.Sleep(time.Second) // Имитация обработки задачи
		results <- job * 2      // Отправляем результат обработки
		fmt.Printf("Воркер %d завершил обработку задачи %d\n", id, job)
	}
}

func main() {
	// Определяем количество доступных ядер процессора
	numCPU := runtime.NumCPU() // 6 ядер на i5 Coffee Lake
	fmt.Printf("Используем %d воркеров (по количеству ядер)\n", numCPU)

	// Создаём каналы для задач и результатов
	jobs := make(chan int, numCPU)
	results := make(chan int, numCPU)

	// Создаём WaitGroup для синхронизации завершения воркеров
	var wg sync.WaitGroup

	// Запускаем пул воркеров
	for i := 1; i <= numCPU; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задачи в канал jobs
	go func() {
		for i := 1; i <= 12; i++ { // 12 задач для обработки
			jobs <- i
		}
		close(jobs) // Закрываем канал после отправки всех задач
	}()

	// Собираем результаты
	go func() {
		wg.Wait()      // Ждём завершения всех воркеров
		close(results) // Закрываем канал результатов
	}()

	// Выводим результаты
	for result := range results {
		fmt.Printf("Результат: %d\n", result)
	}

	fmt.Println("Все задачи завершены.")
}
