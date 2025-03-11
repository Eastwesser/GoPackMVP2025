package main

import "fmt"

/*
	У нас есть рюкзак вместимостью W и набор предметов, каждый из которых имеет вес w[i] и ценность v[i].

	Нужно выбрать набор предметов, чтобы их общий вес не превышал W, а общая ценность была максимальной.
*/

func main() {
	weights := []int{2, 3, 4, 5} // Веса предметов
	values := []int{3, 4, 5, 6}  // Ценности предметов
	capacity := 5                // Вместимость рюкзака

	maxValue := knapsack(weights, values, capacity)
	fmt.Println("Максимальная ценность:", maxValue)
}

// Решение задачи о рюкзаке (O(nW))
func knapsack(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
