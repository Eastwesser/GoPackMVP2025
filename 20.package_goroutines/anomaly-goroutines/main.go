package main

import "fmt"

func cringeGoroutine(firstNum, secondNum int) chan int {
	oniChan := make(chan int)

	go func() {
		defer close(oniChan)

		c := firstNum + secondNum
		oniChan <- c
	}()

	return oniChan
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)
	result := cringeGoroutine(a, b)

	fmt.Println(result)
}
