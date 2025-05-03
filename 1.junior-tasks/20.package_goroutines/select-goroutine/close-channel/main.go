package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch) // закрываем канал, когда записей больше не будет!
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	for v := range ch {
		fmt.Println(v)
	}

	select {
	case message, ok := <-ch:
		if !ok {
			fmt.Println("channel closed") // ЕСЛИ КАНАЛ ЗАКРЫТ, ПОЛУЧАЕМ НУЛЕВОЕ ЗНАЧЕНИЕ ИЛИ ВЫВОДИМ СООБЩЕНИЕ (КАК ТУТ)
		} else {
			fmt.Println("anomaly, we have received:", message)
		}
	}

}
