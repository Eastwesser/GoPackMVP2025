package main

import "fmt"

func main() {
	multiplySlice := func(numbers []int, multiplier int) []int {
		result := make([]int, len(numbers))
		for i, number := range numbers {
			result[i] = multiplier * int(number)
		}
		return result
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	multipliers := multiplySlice(numbers, 10)
	fmt.Println(multipliers)
}
