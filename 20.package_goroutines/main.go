package main

import (
	"fmt"
	"time"
)

func main() {
	oniChan := make(chan string)
	oniChan2 := make(chan int)
	oniChan3 := make(chan int)

	go closer("I can close your Goroutine", oniChan3)
	data3 := <-oniChan3

	go say("Hello World", oniChan, oniChan2)
	go closer("I can close your Goroutine", oniChan3)
	time.Sleep(2 * time.Second)
	data1 := <-oniChan
	data2 := <-oniChan2
	fmt.Println(data1, data2, data3)

	for a := range oniChan3 {
		fmt.Println(a)
	}

	fmt.Println(<-oniChan3)
}

func say(greetings string, oniChan chan string, oniChan2 chan int) {
	fmt.Println("GOROUTINE `SAY` IS RUNNING")
	fmt.Println(greetings)
	oniChan <- "Hello Dondo!"
	oniChan2 <- 1000
	fmt.Println("GOROUTINE `SAY` IS FINISHED")

}

func closer(closeChan string, oniChan3 chan int) {
	defer close(oniChan3)
	fmt.Println("GOROUTINE `CLOSER` IS RUNNING")
	fmt.Println(closeChan)
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		oniChan3 <- 20000
	}
	//oniChan3 <- 20000
	fmt.Println("GOROUTINE `CLOSER` IS FINISHED")
}
