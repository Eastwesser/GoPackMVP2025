package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Кеш для хранения квадратов чисел
type Cache struct {
	mu    sync.Mutex
	items map[int]int
}

// Глобальный кеш
var cache = Cache{
	items: make(map[int]int),
}

// Функция для получения квадрата числа с кешированием
func getSqrt(num int) int {
	// Проверяем кеш
	cache.mu.Lock()
	if val, ok := cache.items[num]; ok {
		cache.mu.Unlock()
		return val
	}
	cache.mu.Unlock()

	// Вычисляем квадрат
	result := num * num

	// Сохраняем в кеш
	cache.mu.Lock()
	cache.items[num] = result
	cache.mu.Unlock()

	return result
}

// Функция для обработки числа (может вызвать панику)
func dealNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Восстановлено после паники при обработке числа %d: %v\n", num, r)
		}
	}()

	// Инициализация генератора случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))
	action := rand.Intn(2) // 0 или 1

	if action == 0 {
		panic(fmt.Sprintf("искусственная паника для числа %d", num))
	}

	// Получаем квадрат числа
	square := getSqrt(num)
	fmt.Printf("Квадрат числа %d = %d\n", num, square)
}

func main() {
	// Настройка количества процессоров
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Используется %d процессоров\n", numCPU)

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Обрабатываем числа от 1 до 10
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go dealNumber(i, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Все числа обработаны")

	// Выводим содержимое кеша
	fmt.Println("\nСодержимое кеша:")
	cache.mu.Lock()
	for num, square := range cache.items {
		fmt.Printf("%d: %d\n", num, square)
	}
	cache.mu.Unlock()
}
