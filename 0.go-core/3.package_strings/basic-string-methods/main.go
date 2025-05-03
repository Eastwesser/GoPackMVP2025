package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello, world!"
	fmt.Println("Original string:", input) // Original string: Hello, world!

	// Разбить строку на символы: []rune
	runes := []rune(input)
	fmt.Printf("Runes: %v\n", runes) // Runes: [72 101 108 108 111 44 32 119 111 114 108 100 33]

	// Разбиение на подстроки
	split := strings.Split(input, ",")
	fmt.Printf("Split by ',' : %v\n", split) // Split by ',' : [Hello  world!]

	splitLimited := strings.SplitN(input, ", ", 2)
	fmt.Printf("Split by ', ' with limit 2: %v\n", splitLimited) // Split by ', ' with limit 2: [Hello world!]

	// Объединение подстрок
	joined := strings.Join(split, " ")
	fmt.Printf("Joined: %s\n", joined) // Joined: Hello  world!

	// Проверка содержимого
	fmt.Printf("Contains 'wol': %v\n", strings.Contains(input, "wol"))   // Contains 'wol': false
	fmt.Printf("HasPrefix 'He': %v\n", strings.HasPrefix(input, "He"))   // HasPrefix 'He': true
	fmt.Printf("HasSuffix 'ld!': %v\n", strings.HasSuffix(input, "ld!")) // HasSuffix 'ld!': true

	// Замена подстроки
	replacedAll := strings.ReplaceAll(input, "Hello", "Hi")
	fmt.Printf("ReplacedAll 'Hello' -> 'Hi': %s\n", replacedAll) // ReplacedAll 'Hello' -> 'Hi': Hi, world!

	replacedLimited := strings.Replace(input, "l", "L", 2)
	fmt.Printf("Replaced first 2 'l' -> 'L': %s\n", replacedLimited) // Replaced first 2 'l' -> 'L': HeLLo, world!

	// Изменение регистра
	lower := strings.ToLower(input)
	fmt.Printf("ToLower: %s\n", lower) // ToLower: hello, world!

	upper := strings.ToUpper(input)
	fmt.Printf("ToUpper: %s\n", upper) // ToUpper: HELLO, WORLD!

	// Удаление пробелов
	trimmed := strings.TrimSpace("   Hello, world!   ")
	fmt.Printf("TrimSpace: '%s'\n", trimmed) // TrimSpace: 'Hello, world!'

	// Поиск символа/подстроки
	firstIndex := strings.Index(input, "o")
	lastIndex := strings.LastIndex(input, "o")
	fmt.Printf("First index of 'o': %d\n", firstIndex) // First index of 'o': 4
	fmt.Printf("Last index of 'o': %d\n", lastIndex)   // Last index of 'o': 8

	// Повторение строки
	repeated := strings.Repeat("Go!", 3)
	fmt.Printf("Repeated 'Go!' 3 times: %s\n", repeated) // Repeated 'Go!' 3 times: Go!Go!Go!

	// Подстрока
	substr := input[1:4]
	fmt.Printf("Substring [1:4]: %s\n", substr) // Substring [1:4]: ell

	// Работа с Builder
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Builder!")
	fmt.Printf("Built string: %s\n", builder.String()) // Built string: Hello, Builder!

}
