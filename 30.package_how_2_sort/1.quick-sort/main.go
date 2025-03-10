package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int // left is 'lesser', right is 'greater'
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	//return quickSort(arr[:len(arr)/2])
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{1236712, 123123, 3125512, 5452, 51235, 512521, 5123132}
	fmt.Println("QUICK SORT:", quickSort(arr))
}
