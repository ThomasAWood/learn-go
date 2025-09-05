package tests

import (
	"learngo/algorithms"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		haystack []int
		needle   int
		expected int
	}{
		{
			name:     "Normal",
			haystack: []int{0, 1, 2, 3, 4, 5},
			needle:   3,
			expected: 3,
		},
		{
			name:     "Not Present",
			haystack: []int{0, 1, 2, 3, 4, 5},
			needle:   72,
			expected: -1,
		},
		{
			name:     "Length One Present",
			haystack: []int{5},
			needle:   5,
			expected: 0,
		},
		{
			name:     "Length One Not Present",
			haystack: []int{5},
			needle:   3,
			expected: -1,
		},
		{
			name:     "First Half",
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			needle:   3,
			expected: 2,
		},
		{
			name:     "Second Half",
			haystack: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			needle:   7,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.BinarySearch(tt.haystack, tt.needle)

			if result != tt.expected {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d",
					tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}
