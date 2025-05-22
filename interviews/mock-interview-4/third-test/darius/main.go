package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func recoverFromPanic(id int) {
	if r := recover(); r != nil {
		fmt.Printf("горутина %d: работа восстановлена после паники: %v\n", id, r)
	}
}
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer recoverFromPanic(id)

	// Генерируем случайное число 0 или 1
	//rand.Seed(time.Now().UnixNano())
	value := rand.Intn(2)

	if value == 0 {
		// Вызываем панику
		panic(fmt.Sprintf("горутина %d: паника!", id))
	} else {
		// Выполняем задание
		fmt.Printf("горутина %d: задание выполнено\n", id)
	}
}

func main() {
	// Устанавливаем количество процессоров
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Используется %d процессоров\n", numCPU)

	// Количество горутин в 10 раз больше чем процессоров
	numWorkers := numCPU * 10
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запускаем горутины
	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Все задания обработаны")
}
