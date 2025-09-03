package main

import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
	testCases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"not a palindrome",
		"",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			isPalindrome(tc)
		}
	}
}

// go test -bench=. -benchmem
