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
			name:     "Normal Test Case",
			haystack: []int{0, 1, 2, 3, 4, 5},
			needle:   3,
			expected: 3,
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
