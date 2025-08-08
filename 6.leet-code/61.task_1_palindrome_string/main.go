package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 125. Valid Palindrome (задача на словарный палиндром)
func isPalindrome(s string) bool {
	// Очищаем строку: оставляем только буквы и цифры, приводим к нижнему регистру
	var cleaned strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			cleaned.WriteRune(unicode.ToLower(ch))
		}
	}
	cleanStr := cleaned.String()

	// Проверяем, является ли очищенная строка палиндромом
	left, right := 0, len(cleanStr)-1
	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	var theWord string

	scan, err := fmt.Scan(&theWord)
	if err != nil {
		return
	}

	fmt.Println(isPalindrome(string(rune(scan))))
}
