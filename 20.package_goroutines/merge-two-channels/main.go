package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	mergedCh := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(mergedCh)
		for num := range ch1 {
			mergedCh <- num
		}
	}()

	go func() {
		for num := range ch2 {
			mergedCh <- num
		}
	}()

	for num := range mergedCh {
		fmt.Println(num)
	}
}
