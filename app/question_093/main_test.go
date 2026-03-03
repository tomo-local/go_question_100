package main

import "testing"

func Test_Add(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "negative",
			a:        -2,
			b:        -4,
			expected: -6,
		},
		{
			name:     "zero",
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			name:     "mixed",
			a:        -1,
			b:        1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
