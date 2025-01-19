package main

import "fmt"

func main() {
	// Создаем три канала для передачи целых чисел
	oniChan1 := make(chan int)    // Канал для первой горутины
	oniChan2 := make(chan int)    // Канал для второй горутины
	mergedChans := make(chan int) // Канал для объединения значений из двух предыдущих каналов

	// Запускаем первую горутину, которая заполняет oniChan1 числами от 1 до 5
	go func() {
		defer close(oniChan1) // Закрываем канал oniChan1 после завершения работы горутины
		for i := 1; i <= 5; i++ {
			oniChan1 <- i // Отправляем числа от 1 до 5 в канал oniChan1
		}
	}()

	// Запускаем вторую горутину, которая заполняет oniChan2 числами от 6 до 10
	go func() {
		defer close(oniChan2) // Закрываем канал oniChan2 после завершения работы горутины
		for i := 6; i <= 10; i++ {
			oniChan2 <- i // Отправляем числа от 6 до 10 в канал oniChan2
		}
	}()

	// Запускаем третью горутину, которая читает из oniChan1 и отправляет значения в mergedChans
	go func() {
		defer close(mergedChans)    // Закрываем mergedChans после завершения работы этой горутины
		for num := range oniChan1 { // Читаем значения из oniChan1 до его закрытия
			mergedChans <- num // Отправляем прочитанные значения в mergedChans
		}
	}()

	// Запускаем четвертую горутину, которая читает из oniChan2 и отправляет значения в mergedChans
	go func() {
		for num := range oniChan2 { // Читаем значения из oniChan2 до его закрытия
			mergedChans <- num // Отправляем прочитанные значения в mergedChans
		}
	}()

	// В основном потоке читаем значения из mergedChans и выводим их на экран
	for num := range mergedChans {
		fmt.Println(num) // Печатаем каждое число, полученное из mergedChans
	}
}
