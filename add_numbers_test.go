package main

import "testing"

func TestAddNumbers(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
	}

	for _, tt := range tests {
		result := add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Expected %d + %d to equal %d, but got %d", tt.a, tt.b, tt.expected, result)
		}
	}
}
