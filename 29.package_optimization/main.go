package main

import (
	"fmt"
	"sync"
)

func syncMapper() {
	// Создаем sync.Map
	var sm sync.Map

	// Добавляем значения в sync.Map
	sm.Store("key1", 1)
	sm.Store("key2", 2)
	sm.Store("key3", 3)

	// Получаем значение по ключу
	if value, ok := sm.Load("key1"); ok {
		fmt.Println("key1:", value)
	}

	// Удаляем значение по ключу
	sm.Delete("key2")

	// Проверяем, существует ли ключ
	if _, ok := sm.Load("key2"); !ok {
		fmt.Println("key2 was deleted")
	}

	// Итерируемся по sync.Map
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true // Продолжаем итерацию
	})
}

func main() {
	syncMapper()
}
