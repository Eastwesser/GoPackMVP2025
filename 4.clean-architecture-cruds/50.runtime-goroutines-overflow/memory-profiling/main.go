package main

import "fmt"

func showMemStats() {}

func main() {
	fmt.Println("Memory before:")
	showMemStats()

	s := make([]int, 10_000_000)

	for i := range s {
		s[i] = i
	}

	fmt.Println("Memory after:")
	showMemStats()
}
