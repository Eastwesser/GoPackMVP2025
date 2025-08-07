package main

import (
	"fmt"
	"regexp"
	"strings"
)

// 125. Valid Palindrome (задача на словарный палиндром)
func isPalindrome(s string) bool {
	var straightOrder string
	var reverseOrder string

	// приводим буквы  в 's' к нижнему регистру
	s = strings.ToLower(s)

	// удаляем в 's' все, кроме букв
	reg := regexp.MustCompile(`[^a-z0-9]`)
	clean := reg.ReplaceAllString(s, "")

	// сохраняем прямой порядок букв
	for i := 0; i < len(clean); i++ {
		straightOrder += string(clean[i])
	}

	// сохраняем реверсивный порядок букв
	for j := len(clean) - 1; j >= 0; j-- {
		reverseOrder += string(clean[j])
	}

	// сравниваем строку с перевернутой
	return straightOrder == reverseOrder
}

func main() {
	var theWord string

	scan, err := fmt.Scan(&theWord)
	if err != nil {
		return
	}

	fmt.Println(isPalindrome(string(rune(scan))))
}
