package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var characterConfig string

// loadCharacterConfig загружает конфигурацию персонажей.
func loadCharacterConfig() {
	characterConfig = "Конфигурация персонажей загружена: Уровни, оружие и артефакты"
	fmt.Println("Загрузка конфигурации персонажей...")
}

func main() {
	var wg sync.WaitGroup

	// Персонажи, которые пытаются загрузить конфигурацию
	characters := []string{"Diluc", "Venti", "Zhongli", "Keqing", "Hu Tao"}

	for _, character := range characters {
		wg.Add(1)
		go func(char string) {
			defer wg.Done()
			once.Do(loadCharacterConfig) // loadCharacterConfig выполнится только один раз
			fmt.Printf("%s: %s\n", char, characterConfig)
		}(character)
	}

	wg.Wait()
}
