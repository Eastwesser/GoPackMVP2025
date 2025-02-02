package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Кринжовая горутина
func cringeGoroutine(ctx context.Context, wg *sync.WaitGroup, dataChan chan<- int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении

	for {
		select {
		case <-ctx.Done(): // Если пришел сигнал завершения
			fmt.Println("Горутина: меня остановили, кринж...")
			return
		default:
			// Спим случайное время (от 100 до 500 мс)
			sleepTime := time.Duration(rand.Intn(400)+100) * time.Millisecond
			fmt.Printf("Горутина: сплю %v...\n", sleepTime)
			time.Sleep(sleepTime)

			// Генерируем случайное число
			num := rand.Intn(100)
			fmt.Printf("Горутина: сгенерировала число %d\n", num)

			// Если число делится на 3, отправляем его в канал
			if num%3 == 0 {
				fmt.Printf("Горутина: число %d делится на 3, отправляю в канал\n", num)
				dataChan <- num
			} else {
				fmt.Printf("Горутина: число %d не делится на 3, пропускаю\n", num)
			}
		}
	}
}

func main() {
	// Создаем контекст с таймаутом для завершения горутины через 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Канал для передачи данных
	dataChan := make(chan int)

	// WaitGroup для ожидания завершения горутины
	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем кринжовую горутину
	go cringeGoroutine(ctx, &wg, dataChan)

	// Горутина для чтения данных из канала
	go func() {
		for num := range dataChan {
			fmt.Printf("Главная функция: получила число %d из канала\n", num)
		}
	}()

	// Ждем завершения горутины
	wg.Wait()

	// Закрываем канал после завершения
	close(dataChan)

	fmt.Println("Программа завершена. Кринж закончился.")
}
