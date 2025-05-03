package main

import "fmt"

/*
4. Скользящее окно
Пример: Анализ урона за последние 5 секунд.

Объяснение: Игрок анализирует урон, чтобы понять, стоит ли атаковать.
*/

func slidingWindow(arr []int, k int) int {
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum := windowSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}

func main() {
	damage := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Max Damage in 5s:", slidingWindow(damage, 3))
}
