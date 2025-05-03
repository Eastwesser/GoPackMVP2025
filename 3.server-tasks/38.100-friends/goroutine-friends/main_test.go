// friends_test.go
package main

import (
	"fmt"
	"testing"
)

func TestAllImplementations(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want string
	}{
		{"1 друг", 1, "друг"},
		{"2 друга", 2, "друга"},
		{"5 друзей", 5, "друзей"},
		{"11 друзей", 11, "друзей"},
		{"21 друг", 21, "друг"},
		{"22 друга", 22, "друга"},
		{"25 друзей", 25, "друзей"},
		{"101 друг", 101, "друг"},
		{"111 друзей", 111, "друзей"},
	}

	implementations := []struct {
		name string
		fn   func(int) string
	}{
		{"getFriendWord", getFriendWord},
		{"processFriendsConcurrent", func(n int) string { return processFriendsConcurrent(n)[n-1] }},
	}

	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					if got := impl.fn(tc.n); got != tc.want {
						t.Errorf("%s(%d) = %v, want %v", impl.name, tc.n, got, tc.want)
					}
				})
			}
		})
	}
}

func BenchmarkAllImplementations(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(int) []string
	}{
		{"Concurrent", processFriendsConcurrent},
	}

	sizes := []int{100, 1000, 10000, 100000}

	for _, bench := range benchmarks {
		for _, size := range sizes {
			b.Run(fmt.Sprintf("%s-%d", bench.name, size), func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					bench.fn(size)
				}
			})
		}
	}
}
