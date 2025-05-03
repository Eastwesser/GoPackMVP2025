package main

import (
	"fmt"
	"time"
)

// функция Эола с каналом Lawrence
func Eula(Lawrence chan string) {
	time.Sleep(3 * time.Second) // Эола спит дольше, 3 часа
	Lawrence <- "VENGEANCE WILL BE MINE!"
	fmt.Println("Eula wins!")
}

// функция Эола с каналом Lawrence
func Mavuika(PyroArchon chan string) {
	time.Sleep(2 * time.Second) // Мавуика спит меньше Эолы, 2 часа
	PyroArchon <- "FOR NATLAN!"
	fmt.Println("Mavuika wins!")
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go Eula(ch1)
	go Mavuika(ch2)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("Канал пуст, продолжаем работу") // дефолт можно не писать, но лучше писать, помогает избежать блокировки, если каналы пустые
	}

	fmt.Println("Genshin Impact")
}
