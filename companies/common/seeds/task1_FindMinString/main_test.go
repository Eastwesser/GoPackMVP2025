package main

import (
	"testing"
)

func TestFindMinString(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    string
		wantErr bool
	}{
		{"1", []string{"hesoyam", "yesgaa", "uzumumw"}, "hesoyam", false},
		{"2", []string{"Hmm", "oh wow", "no way bro!"}, "Hmm", false},
		{"3", []string{""}, "", false},
		{"3", []string{}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMinString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMinString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindMinString() = %v, want %v", got, tt.want)
			}
		})
	}
}
