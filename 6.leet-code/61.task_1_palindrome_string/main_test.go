package main

import "testing"

// BENCHMARK

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("racecar")
	}
}
