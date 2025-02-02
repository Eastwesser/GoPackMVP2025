package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func playerAction(playerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	actionTime := rand.Intn(5) + 1 // Время действия от 1 до 5 секунд
	time.Sleep(time.Duration(actionTime) * time.Second)
	fmt.Printf("%s завершил действие\n", playerName)
}

func dragonSpawn(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Second) // Дракон появляется через 10 секунд
	fmt.Println("Дракон появился!")
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Для случайных значений

	var wg sync.WaitGroup

	players := []string{"Игрок 1", "Игрок 2", "Игрок 3", "Игрок 4", "Игрок 5"}

	for _, player := range players {
		wg.Add(1)
		go playerAction(player, &wg)
	}

	wg.Add(1)
	go dragonSpawn(&wg)

	wg.Wait()
	fmt.Println("Матч продолжается после убийства дракона!")
}
