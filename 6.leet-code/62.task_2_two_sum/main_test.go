package main

import (
	"reflect"
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

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Five", []int{1, 2, 3}, 5, []int{1, 2}},
		{"Ten", []int{2, 6, 5, 9, 4, 3, 7}, 10, []int{1, 4}},
		{"Hundred", []int{25, 64, 51, 92, 48, 36, 73}, 100, []int{1, 5}},
	}

	for _, tt := range tests {
		result := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(result, tt.want) {
			if len(result) != 2 || result[0] != tt.target || result[1] != tt.target {
				t.Error("twoSum", tt.nums, tt.target, "expected", tt.want, "got", result)
			}
		}
	}
}
