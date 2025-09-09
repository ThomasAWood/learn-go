package tests

import (
	"learngo/algorithms"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal Test Case",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty List",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Length One",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "All the same",
			input:    []int{2, 2, 2, 2, 2, 2},
			expected: []int{2, 2, 2, 2, 2, 2},
		},
		{
			name:     "Duplicate Numbers",
			input:    []int{3, 5, 3, 5, 3, 2, 1, 9},
			expected: []int{1, 2, 3, 3, 3, 5, 5, 9},
		},
	}

	// Run all tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.BubbleSort(tt.input)

			for i, elem := range result {
				if tt.expected[i] != elem {
					t.Errorf("BubbleSort(%v) \n Expected: %v \n Got: %v",
						tt.input, tt.expected, result)
				}
			}
		})
	}
}
