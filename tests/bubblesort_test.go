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
