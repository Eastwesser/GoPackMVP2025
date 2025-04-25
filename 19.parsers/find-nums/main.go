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
	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	// Regular expression to find numbers in the text
	re := regexp.MustCompile(`(\d[\d\s]*\d|\d)`)

	total := 0

	// Process each line
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		if len(matches) == 0 {
			continue
		}

		// Process each number found in the line
		for _, match := range matches {
			cleanNum := strings.ReplaceAll(match, " ", "")
			num, err := strconv.Atoi(cleanNum)
			if err == nil {
				total += num
			}
		}
	}

	fmt.Printf("Total sum of all numbers: %d\n", total)
}
