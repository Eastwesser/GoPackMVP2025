package main

import "fmt"

func quickMath(a, b int) chan int {
	oniChan := make(chan int)

	go func() {
		defer close(oniChan)
		sum := a + b
		oniChan <- sum
		fmt.Println("Сумма посчитана. Горутина передала значение в канал!")
	}()

	return oniChan
}

func main() {
	a, b := 2, 2
	oniChan := quickMath(a, b)
	fmt.Println(<-oniChan)
	// считывание результата из канала
	result := <-quickMath(a, b)
	fmt.Println("Сумма", result)
}
