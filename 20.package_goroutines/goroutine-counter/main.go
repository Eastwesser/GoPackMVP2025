package main

import (
	"fmt"
	"time"
)

func numberGenerator(n int) chan int {
	//fmt.Println("Goroutine `generateData` is running")

	data := make(chan int)

	go func() {
		defer close(data)

		for i := 1; i <= n; i++ {
			data <- i
		}

	}()

	return data
}

func main() {
	data := numberGenerator(30)

	go func() {
		time.Sleep(1 * time.Second)
		close(data)
	}()

	for num := range data {
		fmt.Println(num)
	}

}
