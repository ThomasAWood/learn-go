package tests

import (
	"learngo/algorithms"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		haystack []int
		needle   int
		expected int
	}{
		{
			name:     "Normal Test Case",
			haystack: []int{0, 1, 2, 3, 4, 5},
			needle:   3,
			expected: 3,
		},
		{
			name:     "Not Present",
			haystack: []int{3, 4, 5, 6, 7, 8},
			needle:   6,
			expected: 3,
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
			needle:   9,
			expected: -1,
		},
		{
			name:     "Empty List",
			haystack: []int{},
			needle:   9,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.LinearSearch(tt.haystack, tt.needle)

			if result != tt.expected {
				t.Errorf("LinearSearch(%v, %d) = %d, want %d",
					tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}
