package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	text, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\d{1,3}(?:\s?\d{3})*`)
	total := 0

	for _, match := range re.FindAllString(string(text), -1) {
		if num, err := strconv.Atoi(strings.ReplaceAll(match, " ", "")); err == nil {
			total += num
		}
	}

	fmt.Printf("Total sum of all numbers: %d\n", total)
}
