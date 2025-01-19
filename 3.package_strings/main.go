package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello, world!"
	fmt.Println("Original string:", input)

	// Разбить строку на символы: []rune
	runes := []rune(input)
	fmt.Printf("Runes: %v\n", runes)

	// Разбиение на подстроки
	split := strings.Split(input, ",")
	fmt.Printf("Split by ',' : %v\n", split)

	splitLimited := strings.SplitN(input, ", ", 2)
	fmt.Printf("Split by ', ' with limit 2: %v\n", splitLimited)

	// Объединение подстрок
	joined := strings.Join(split, " ")
	fmt.Printf("Joined: %s\n", joined)

	// Проверка содержимого
	fmt.Printf("Contains 'wol': %v\n", strings.Contains(input, "wol"))
	fmt.Printf("HasPrefix 'He': %v\n", strings.HasPrefix(input, "He"))
	fmt.Printf("HasSuffix 'ld!': %v\n", strings.HasSuffix(input, "ld!"))

	// Замена подстроки
	replacedAll := strings.ReplaceAll(input, "Hello", "Hi")
	fmt.Printf("ReplacedAll 'Hello' -> 'Hi': %s\n", replacedAll)

	replacedLimited := strings.Replace(input, "l", "L", 2)
	fmt.Printf("Replaced first 2 'l' -> 'L': %s\n", replacedLimited)

	// Изменение регистра
	lower := strings.ToLower(input)
	fmt.Printf("ToLower: %s\n", lower)

	upper := strings.ToUpper(input)
	fmt.Printf("ToUpper: %s\n", upper)

	// Удаление пробелов
	trimmed := strings.TrimSpace("   Hello, world!   ")
	fmt.Printf("TrimSpace: '%s'\n", trimmed)

	// Поиск символа/подстроки
	firstIndex := strings.Index(input, "o")
	lastIndex := strings.LastIndex(input, "o")
	fmt.Printf("First index of 'o': %d\n", firstIndex)
	fmt.Printf("Last index of 'o': %d\n", lastIndex)

	// Повторение строки
	repeated := strings.Repeat("Go!", 3)
	fmt.Printf("Repeated 'Go!' 3 times: %s\n", repeated)

	// Подстрока
	substr := input[1:4]
	fmt.Printf("Substring [1:4]: %s\n", substr)

	// Работа с Builder
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Builder!")
	fmt.Printf("Built string: %s\n", builder.String())
}
