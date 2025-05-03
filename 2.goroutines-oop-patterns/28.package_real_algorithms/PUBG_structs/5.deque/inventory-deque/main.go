package main

import (
	"container/list"
	"fmt"
)

func main() {
	inventory := list.New() // Дек для инвентаря

	// Добавляем предметы в инвентарь
	inventory.PushBack("Ammo")    // Обычный предмет
	inventory.PushFront("Medkit") // Срочный предмет
	inventory.PushBack("Grenade") // Обычный предмет

	// Используем предметы
	for inventory.Len() > 0 {
		// Сначала используем срочные предметы из начала
		frontItem := inventory.Front()
		fmt.Println("Using Front Item:", frontItem.Value)
		inventory.Remove(frontItem)

		// Затем используем обычные предметы из конца
		if inventory.Len() > 0 {
			backItem := inventory.Back()
			fmt.Println("Using Back Item:", backItem.Value)
			inventory.Remove(backItem)
		}
	}

	fmt.Println("Inventory is empty.")
}
