package main

import (
	"fmt"
	"sync"
)

func main() {
	const maxGoroutines = 3                       // Максимальное количество одновременно выполняющихся горутин
	oniChan := make(chan struct{}, maxGoroutines) // Создаем буферизированный канал с размером maxGoroutines
	var wgPlanner sync.WaitGroup                  // Создаем WaitGroup для ожидания завершения всех горутин

	for i := 1; i <= 5; i++ {
		wgPlanner.Add(1)      // Увеличиваем счетчик WaitGroup на 1
		oniChan <- struct{}{} // Отправляем пустую структуру в канал

		go func(num int) {
			defer wgPlanner.Done()       // Уменьшаем счетчик WaitGroup при завершении работы горутины
			defer func() { <-oniChan }() // Освобождаем место в канале после завершения работы

			fmt.Println("Горутина", num, "выполняется") // Выводим номер выполняемой горутины
		}(i) // Передаем номер итерации в горутину
	}

	wgPlanner.Wait() // Ожидаем завершения всех горутин

	fmt.Println("Все горутины завершены.") // Сообщение о завершении всех горутин
}
