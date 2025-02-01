package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	В этом коде реализован Worker Pool — пул воркеров (горутин), которые обрабатывают числа.
	Используется sync.WaitGroup, context.Context и runtime.NumCPU().
*/

func main() {
	baseKnowledge()
	workerPool() // Основная функция, которая создаёт пул воркеров и управляет потоками данных.
}

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	toDo := context.TODO()
	fmt.Println(toDo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	cancel()
	fmt.Println(withCancel.Err())

	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println(withTimeout.Done())
}

func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*20) // Создаём контекст с таймаутом на 20 мс.
	defer cancel()                                                               // cancel() вызывается при выходе из функции, чтобы завершить все воркеры.

	wg := &sync.WaitGroup{} // Создаём WaitGroup — он будет ждать завершения всех воркеров.

	/*
		Создаём два канала:
			numbersToProcess 📦 — отправляем числа, которые нужно обработать.
			processedNumbers ✅ — готовые результаты.
	*/
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ { //  Запускаем воркеров (столько же, сколько ядер процессора).
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*
				Каждый воркер:
					Берёт число из numbersToProcess.
						Вычисляет его квадрат.
						Отправляет результат в processedNumbers.
						Работает, пока канал не закроется или не истечёт контекст.
			*/
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		// Заполняем канал задачами (0...999).
		for i := 0; i < 1000; i++ {
			if i == 500 {
				cancel()
			}
			numbersToProcess <- i
		}
		close(numbersToProcess) // Закрываем канал, когда отправили все числа.
	}()

	go func() {
		wg.Wait()               // Ожидаем завершения всех воркеров.
		close(processedNumbers) //  Закрываем processedNumbers, когда все воркеры отработали.
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++ //  Читаем и печатаем все результаты.
		fmt.Println(resultValue)
	}

	fmt.Println(counter) // Считаем, сколько чисел обработано.
}

// Функция воркера — получает числа, обрабатывает их и отправляет результат.
func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			// Если контекст отменён, воркер завершает работу.
			return
		case value, ok := <-toProcess:
			// Читаем число из канала.
			if !ok {
				//  Если канал закрыт, выходим.
				return
			}
			time.Sleep(time.Millisecond) //  Имитация работы (Sleep на 1 мс).
			processed <- value * value   // Вычисляем квадрат числа и отправляем в processedNumbers.
		}
	}
}

/*
	Создаётся контекст с таймаутом на 20 мс.
	Запускается пул воркеров (по числу ядер процессора).
	В канал numbersToProcess отправляются числа 0...999.
	Воркеры читают числа, считают их квадраты и отправляют в processedNumbers.
	Когда все числа отправлены, канал numbersToProcess закрывается.
	Когда все воркеры отработали, processedNumbers тоже закрывается.
	Выводятся обработанные числа.
*/
