package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Artifact представляет информацию об артефакте.
type Artifact struct {
	Type   string // Тип артефакта (например, "Цветок жизни", "Перо смерти")
	Rarity int    // Редкость артефакта (например, 3, 4, 5)
}

// mergeChan объединяет несколько каналов в один.
func mergeChan(ch ...chan Artifact) chan Artifact {
	var wg sync.WaitGroup
	mergedChan := make(chan Artifact)

	// Функция для чтения из одного канала и записи в объединенный канал.
	readAndWrite := func(c <-chan Artifact) {
		defer wg.Done()
		for artifact := range c {
			mergedChan <- artifact
		}
	}

	wg.Add(len(ch))
	for _, c := range ch {
		go readAndWrite(c)
	}

	// Закрываем объединенный канал после завершения всех горутин.
	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}

func main() {
	// Создаем каналы для передачи артефактов.
	ch1 := make(chan Artifact)
	ch2 := make(chan Artifact)
	ch3 := make(chan Artifact)
	ch4 := make(chan Artifact)

	// Объединяем каналы в один.
	mergedChan := mergeChan(ch1, ch2, ch3, ch4)

	// Запускаем горутины для записи артефактов в каналы.
	go writeArtifacts(ch1, "Цветок жизни")
	go writeArtifacts(ch2, "Перо смерти")
	go writeArtifacts(ch3, "Пески времени")
	go writeArtifacts(ch4, "Кубок пространства")

	// Читаем и выводим артефакты из объединенного канала.
	counter := 0
	for artifact := range mergedChan {
		counter++
		fmt.Printf("Артефакт %d: Тип = %s, Редкость = %d\n", counter, artifact.Type, artifact.Rarity)
	}

	fmt.Println("Все артефакты обработаны.")
}

// RandArtifacts генерирует случайные артефакты.
func RandArtifacts(length int) []Artifact {
	var artifacts []Artifact
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		rarity := rand.Intn(3) + 3 // Редкость от 3 до 5
		artifacts = append(artifacts, Artifact{
			Type:   "Случайный артефакт",
			Rarity: rarity,
		})
	}
	return artifacts
}

// writeArtifacts записывает артефакты в канал.
func writeArtifacts(ch chan<- Artifact, artifactType string) {
	defer close(ch)
	for _, artifact := range RandArtifacts(10) { // Генерируем 10 артефактов
		artifact.Type = artifactType
		ch <- artifact
	}
}
