package main

import (
	"fmt"
	"time"
)

func main() {
	go say()
	time.Sleep(2 * time.Second)
}

func say() {
	fmt.Println("Hello World")
}
