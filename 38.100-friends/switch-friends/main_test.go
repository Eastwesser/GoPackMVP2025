package main

import "testing"

func TestGetFriendWord(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFriendWord(tt.n); got != tt.want {
				t.Errorf("getFriendWord(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkGetFriendWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFriendWord(i % 1000)
	}
}

func BenchmarkProcessFriendsSequential100(b *testing.B) {
	benchmarkProcessFriends(100, b)
}

func BenchmarkProcessFriendsSequential10000(b *testing.B) {
	benchmarkProcessFriends(10000, b)
}

func benchmarkProcessFriends(count int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFriendsSequential(count)
	}
}
