package main

import "fmt"

/*
7. Бэктрекинг
Пример: Поиск оптимального маршрута для сбора лута.

Объяснение: Игрок ищет маршрут, чтобы собрать максимальный лут.
*/

func backtrack(loot []string, path []string, result *[][]string) {
	if len(path) == len(loot) {
		*result = append(*result, append([]string{}, path...))
		return
	}
	for _, item := range loot {
		if !contains(path, item) {
			path = append(path, item)
			backtrack(loot, path, result)
			path = path[:len(path)-1]
		}
	}
}

func contains(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func main() {
	loot := []string{"Ammo", "Medkit", "Grenade"}
	var result [][]string
	backtrack(loot, []string{}, &result)
	fmt.Println("Possible Loot Routes:", result)
}
