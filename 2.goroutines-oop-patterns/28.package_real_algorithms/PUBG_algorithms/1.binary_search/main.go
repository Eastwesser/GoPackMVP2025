package main

import "fmt"

/*
1. Бинарный поиск
Пример: Поиск оружия в отсортированном списке.

Объяснение: Игрок ищет оружие в инвентаре.
*/

func binarySearch(arr []string, target string) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	weapons := []string{"AKM", "M416", "SKS", "UMP45"}
	index := binarySearch(weapons, "M416")
	fmt.Println("Weapon found at index:", index)
}
