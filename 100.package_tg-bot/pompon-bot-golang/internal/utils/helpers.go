package utils

import (
	"log"
	"regexp"
	"strings"
)

// CheckError логирует ошибку, если она не nil.
func CheckError(err error) {
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

// CheckFatalError логирует ошибку и завершает программу, если она не nil.
func CheckFatalError(err error) {
	if err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
}

// ToSnakeCase преобразует строку в snake_case.
func ToSnakeCase(str string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// IsEmpty проверяет, является ли строка пустой или состоит только из пробелов.
func IsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
