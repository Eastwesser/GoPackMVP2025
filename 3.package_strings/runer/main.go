package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Функция переворачивает строку, представленную в виде среза рун.
func reverseRunes(runes []rune) []rune {
	n := len(runes)
	reversed := make([]rune, n) // Создаём новый срез для результата.
	for i, r := range runes {
		reversed[n-i-1] = r // Размещаем символы в обратном порядке.
	}
	return reversed
}

// Функция заменяет все гласные символы в строке на '*'.
func replaceVowelsWithAsterisk(text string) string {
	runes := []rune(text)     // Преобразуем строку в руны.
	vowels := "aeiouAEIOU"    // Гласные символы.
	for i, r := range runes { // Проходим по всем рунам.
		if strings.ContainsRune(vowels, r) {
			runes[i] = '*' // Заменяем гласную на '*'.
		}
	}
	return string(runes) // Преобразуем руны обратно в строку.
}

func main() {
	text := "Hello, 世界! 🌍"

	// Разделяем строку на руны.
	runes := []rune(text)
	fmt.Printf("Runes: %q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界' '!' ' ' '🌍']

	// Подсчитываем количество рун.
	fmt.Println("Number of runes:", len(runes)) // Количество рун: 13.

	// Проверяем, является ли каждый символ буквой, цифрой или другим символом.
	for _, r := range runes {
		if unicode.IsLetter(r) {
			fmt.Printf("Letter: %q\n", r) // Вывод букв.
		} else if unicode.IsDigit(r) {
			fmt.Printf("Digit: %q\n", r) // Вывод цифр.
		} else {
			fmt.Printf("Other: %q\n", r) // Остальные символы.
		}
	}

	// Переворачиваем строку.
	reversed := reverseRunes(runes)
	fmt.Printf("Reversed: %q\n", reversed) // ['🌍' ' ' '!' '界' '世' ' ' ',' 'o' 'l' 'l' 'e' 'H'].

	// Заменяем гласные символы на '*'.
	replaced := replaceVowelsWithAsterisk(text)
	fmt.Printf("Replaced vowels: %s\n", replaced) // H*ll*, 世*🌍
}
