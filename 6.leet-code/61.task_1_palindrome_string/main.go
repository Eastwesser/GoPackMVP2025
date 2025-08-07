package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	var theWord string

	scan, err := fmt.Scan(&theWord)
	if err != nil {
		fmt.Println(err)
	}
	isPalindrome(theWord)

	fmt.Println(scan)
}
