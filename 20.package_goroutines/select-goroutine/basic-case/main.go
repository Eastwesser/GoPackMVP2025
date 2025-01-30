package main

import (
	"fmt"
	"time"
)

func worker(oniChan chan string) {
	time.Sleep(4 * time.Second) // если воркер спит меньше чем передано ниже в 21 строке, то данные передаются в канал
	oniChan <- "Данные из воркера передаются в канал oniChan"
}

func main() {
	ch := make(chan string)

	go worker(ch) // это запуск горутины

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(6 * time.Second): // таймаут это некий дэдлайн по времени
		fmt.Println("timeout - no data received")
	}

	fmt.Println("Hello World")
}
