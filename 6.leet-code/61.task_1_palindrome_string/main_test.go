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

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Normal", "racecar", true},
		{"Camel", "RaDaR", true},
		{"Not palindrome", "dinosaur", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
