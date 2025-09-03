package main

import (
	"testing"
)

func BenchmarkTwoSum(b *testing.B) {
	testCases := []int{
		1, 2, 3, 64, 36, 56,
	}
	target := 100

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		twoSum(testCases, target)
	}
}

// go test -bench=. -benchmem
