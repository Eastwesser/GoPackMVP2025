package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{1, 5, 10, 25} // Номиналы монет
	amount := 63                  // Сумма

	minCoins := coinChange(coins, amount)
	fmt.Println("Минимальное количество монет:", minCoins)
}

// Жадный алгоритм для задачи о минимальном количестве монет
func coinChange(coins []int, amount int) int {
	// Сортируем монеты по убыванию
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	count := 0
	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			count++
		}
	}

	if amount == 0 {
		return count
	}
	return -1 // Если сумму нельзя набрать
}