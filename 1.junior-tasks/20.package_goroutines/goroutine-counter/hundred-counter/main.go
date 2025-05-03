package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(start int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second) // Задержка для визуализации
	}
}

func main() {
	var b int
	fmt.Println("Введите число от 1 до 100:")
	fmt.Scanln(&b)

	if b < 1 || b > 100 {
		fmt.Println("Число должно быть от 1 до 100.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go countdown(b, &wg)

	wg.Wait()
	fmt.Println("Отсчёт завершен!")
}
