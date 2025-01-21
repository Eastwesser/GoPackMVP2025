package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}

			if !swapped {
				break
			}
		}
	}
	return arr
}

func main() {
	newArray := []int{125367, 12329, 8736, 21}
	bubby := bubbleSort(newArray)
	fmt.Println(bubby)
}
