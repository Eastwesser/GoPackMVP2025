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
