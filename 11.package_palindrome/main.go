package main

import (
	"fmt"
	"strings"
)

// Функция проверяет, является ли строка палиндромом.
func isPalindrome(str string) bool {
	str = strings.ToLower(str)             // Приводим строку к нижнему регистру.
	str = strings.ReplaceAll(str, " ", "") // Убираем пробелы из строки.

	// Сравниваем символы с начала и конца строки.
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] { // Если символы не совпадают, возвращаем false.
			return false
		}
	}
	return true // Если цикл завершился без ошибок, строка — палиндром.
}

func main() {
	var newWord string
	fmt.Print("Enter a word: ") // Запрашиваем ввод у пользователя.
	fmt.Scanln(&newWord)        // Читаем введённую строку.

	// Проверяем, является ли введённая строка палиндромом.
	if isPalindrome(newWord) {
		fmt.Println("It's a palindrome!") // Если строка палиндром.
	} else {
		fmt.Println("Not a palindrome.") // Если не палиндром.
	}
}
