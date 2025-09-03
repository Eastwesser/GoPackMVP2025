package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return []int{}
}

func main() {
	var numbers int
	fmt.Println("Enter the number of elements: ")
	fmt.Scan(&numbers)

	intSlice := make([]int, numbers)
	fmt.Printf("Enter %d elements of the array:\n", numbers)
	for i := 0; i < numbers; i++ {
		fmt.Scan(&intSlice[i]) // почему именно так? мы распаковываем по индексу каждый элемент?
	}

	var sumNum int
	fmt.Println("Enter the target (sum of two numbers in the array):")
	fmt.Scan(&sumNum)

	result := twoSum(intSlice, sumNum)
	fmt.Println("Their indexes are:", result)
}
