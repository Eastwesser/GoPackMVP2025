package main

import "fmt"

func processFriendsOptimized(max int) []string {
	result := make([]string, max)
	for i := range result {
		n := i + 1
		if n%100 >= 11 && n%100 <= 14 {
			result[i] = "друзей"
			continue
		}
		switch n % 10 {
		case 1:
			result[i] = "друг"
		case 2, 3, 4:
			result[i] = "друга"
		default:
			result[i] = "друзей"
		}
	}
	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("У меня %d %s!\n", i, processFriendsOptimized(i))
	}
}
