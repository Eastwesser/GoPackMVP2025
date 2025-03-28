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
	results := processFriendsOptimized(100)
	for i, word := range results {
		fmt.Printf("У меня %d %s!\n", i+1, word)
	}
}
