package main

import (
	"fmt"
	"time"
)

func main() {
	oniChan := make(chan int)

	go say("Hello World", oniChan)
	time.Sleep(2 * time.Second)
	fmt.Println(<-oniChan)
}

func say(greetings string, oniChan chan int) {
	fmt.Println(greetings)
	oniChan <- 7
}
