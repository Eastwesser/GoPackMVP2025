package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "   Hello, Golang!   "
	fmt.Println("Original string:", input)

	// Trim whitespace
	trimmed := strings.TrimSpace(input)
	fmt.Println("Trimmed string:", trimmed)

	// Convert to uppercase
	upper := strings.ToUpper(trimmed)
	fmt.Println("Uppercase string:", upper)

	// Replace a substring
	replaced := strings.ReplaceAll(upper, "GOLANG", "WORLD")
	fmt.Println("Replaced string:", replaced)

	// Split string into words
	words := strings.Fields(trimmed)
	fmt.Println("Words in the string:", words)

	// Join words back into a string
	joined := strings.Join(words, ", ")
	fmt.Println("Joined words:", joined)
}
