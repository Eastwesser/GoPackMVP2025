package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// горутина любви, выдает текст "я люблю тебя" 100 раз в цикле
func loveGoroutine(ctx context.Context, wg *sync.WaitGroup, loveChan chan<- string) {
	defer wg.Done()       // Уменьшаем счетчик WaitGroup при завершении
	defer close(loveChan) // Закрываем канал dataChan сразу, как горутина отработает

	var count int
	for {
		select {
		case <-ctx.Done(): // Если пришел сигнал завершения
			fmt.Println("Горутина: меня остановили")
			return
		default:
			if count >= 100 {
				return
			}
			// Генерируем любовное послание
			loveChan <- "Я люблю тебя"
			//fmt.Print("Ура, любовь пошла!")
			count++
		}
	}
}

func main() {
	// Создаем контекст с таймаутом для завершения горутины через 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Канал для передачи данных
	loveChan := make(chan string)

	// WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup
	wg.Add(1)
	//wg.Add(2)

	// Запускаем любовную горутину
	go loveGoroutine(ctx, &wg, loveChan)

	// Горутина-читатель для чтения данных из канала
	go func() {
		//defer wg.Done() // если WaitGroup равно 2 (при 1 будет отрицалетьный счет и паника)
		for msg := range loveChan {
			fmt.Printf("Главная функция получила из канала послание для Ники:\n %s \n", msg)
		}
	}()

	// Ждем завершения горутины
	wg.Wait()
	fmt.Println("Программа завершена. Цикл любви закончился.")
}
