package main

import (
	"fmt"
)

func main() {
	oniChan := make(chan string)
	oniChan2 := make(chan int)

	go say("Hello World", oniChan, oniChan2)
	//time.Sleep(2 * time.Second)
	data1 := <-oniChan
	data2 := <-oniChan2
	fmt.Println(data1, data2)
}

func say(greetings string, oniChan chan string, oniChan2 chan int) {
	fmt.Println(greetings)
	oniChan <- "HI HI"
	oniChan2 <- 1000
}
