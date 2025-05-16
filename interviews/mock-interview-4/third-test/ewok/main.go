package main

import (
	"fmt"
	"sync"
)

// ewok выполняет операции в зависимости от значения n
func ewok(n int) int {
	if n == 0 {
		panic("n не может быть равен 0")
	}
	return n * n
}

// safeEwok оборачивает вызов ewok с обработкой паники
func safeEwok(n int) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = "Произошла какая-то ошибка"
		}
	}()

	res := ewok(n)
	return fmt.Sprintf("%d", res)
}

func main() {
	var wg sync.WaitGroup

	// Запускаем горутины для значений от -50 до 50
	for i := -50; i <= 50; i++ {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			res := safeEwok(n)
			fmt.Printf("Результат для %d: %s\n", n, res)
		}(i)

	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}
