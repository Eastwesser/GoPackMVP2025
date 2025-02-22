package main

import (
	"fmt"
	"sync"
)

// MonsterCounter представляет счетчик побежденных монстров.
type MonsterCounter struct {
	mu      sync.Mutex // Мьютекс для синхронизации доступа к счетчику
	counter int        // Счетчик монстров
}

// Increment увеличивает счетчик побежденных монстров.
func (mc *MonsterCounter) Increment() {
	mc.mu.Lock() // Блокируем доступ к ресурсу
	mc.counter++
	mc.mu.Unlock() // Освобождаем доступ
}

func main() {
	var wg sync.WaitGroup
	mc := &MonsterCounter{} // Создаем счетчик монстров

	// Персонажи, которые будут побеждать монстров
	characters := []string{"Diluc", "Venti", "Zhongli", "Keqing"}

	// Количество монстров, которых нужно победить
	monstersToDefeat := 1000

	// Запускаем горутины для каждого персонажа
	for _, character := range characters {
		wg.Add(1)
		go func(char string) {
			defer wg.Done()
			for i := 0; i < monstersToDefeat/len(characters); i++ {
				mc.Increment() // Персонаж побеждает монстра
				fmt.Printf("%s победил монстра! Всего побеждено: %d\n", char, mc.counter)
			}
		}(character)
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Printf("Все монстры побеждены! Итоговый счетчик: %d\n", mc.counter)
}
