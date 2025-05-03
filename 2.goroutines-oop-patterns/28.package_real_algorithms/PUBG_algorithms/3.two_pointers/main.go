package main

import "fmt"

/*
3. Два указателя
Пример: Поиск двух ближайших противников.

Объяснение: Игрок ищет двух ближайших противников для атаки.
*/

func twoPointers(arr []int, target int) (int, int) {
	left, right := 0, len(arr)-1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return arr[left], arr[right]
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

func main() {
	distances := []int{10, 20, 30, 40, 50}
	a, b := twoPointers(distances, 60)
	fmt.Println("Closest Enemies:", a, b)
}
