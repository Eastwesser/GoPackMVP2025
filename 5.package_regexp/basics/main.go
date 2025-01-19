package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Example: Check if a string matches an email pattern
	email := "example@domain.com"
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, _ := regexp.MatchString(emailRegex, email)
	fmt.Printf("Does '%s' match the email regex? %v\n", email, matched)

	// Example: Replace phone numbers with "[REDACTED]"
	text := "Call me at 123-456-7890 or 987-654-3210."
	phoneRegex := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)

	censored := phoneRegex.ReplaceAllString(text, "[REDACTED]")
	fmt.Println("Censored text:", censored)

	// Example: Extract all words starting with "Go"
	sentence := "Go is a programming language. I love GoLang and Go's simplicity."
	wordRegex := regexp.MustCompile(`\bGo[a-zA-Z]*\b`)

	words := wordRegex.FindAllString(sentence, -1)
	fmt.Println("Words starting with 'Go':", words)
}
