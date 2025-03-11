package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"колодка", "Лох", "ка", "ло", "док"} // Список слов
	target := "лодка"                                      // Целевое слово

	if canBuildWord(target, words) {
		fmt.Println("Слово можно составить!")
	} else {
		fmt.Println("Слово нельзя составить.")
	}
}

// Проверка, можно ли составить слово из частей
func canBuildWord(target string, words []string) bool {
	if target == "" {
		return true
	}

	for _, word := range words {
		if strings.HasPrefix(target, word) {
			if canBuildWord(target[len(word):], words) {
				return true
			}
		}
	}

	return false
}
