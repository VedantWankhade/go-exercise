package random

import (
	"fmt"
	"testing"
)

func TestRandomStringWithLength(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Length 0",
			args: 0,
			want: 0,
		},
		{
			name: "Length 1",
			args: 1,
			want: 1,
		},
		{
			name: "Length 2",
			args: 2,
			want: 2,
		},
		{
			name: "Length 3",
			args: 3,
			want: 3,
		},
		{
			name: "Length 12",
			args: 12,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomStringWithLength(tt.args)
			fmt.Println("=> " + got)
			if len(got) != tt.want {
				t.Errorf("RandomStringWithLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
		},
		{
			name: "Test 2",
		},
		{
			name: "Test 3",
		},
		{
			name: "Test 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(22)
			fmt.Println("=> " + got)
			if len(got) == 0 {
				t.Errorf("RandomString() = %v, want %v", got, "not zero")
			}
		})
	}
}
