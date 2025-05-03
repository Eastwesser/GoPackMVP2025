package main

import "fmt"

/*
	Фермер хочет разделить поле на участки для посадки разных культур.
	У него есть ограничения на площадь и ресурсы.
	Нужно максимизировать прибыль.
*/

func main() {
	areas := []int{1, 2, 3, 4}   // Площади участков
	profits := []int{2, 3, 4, 5} // Прибыль с участков
	totalArea := 5               // Общая площадь

	maxProfit := maximizeProfit(areas, profits, totalArea)
	fmt.Println("Максимальная прибыль:", maxProfit)
}

// Решение задачи максимизации прибыли (O(nW))
func maximizeProfit(areas, profits []int, totalArea int) int {
	n := len(areas)
	dp := make([]int, totalArea+1)

	for i := 0; i < n; i++ {
		for w := totalArea; w >= areas[i]; w-- {
			dp[w] = max(dp[w], dp[w-areas[i]]+profits[i])
		}
	}

	return dp[totalArea]
}
