package preference

import (
	"testing"
)

func TestGetKthLargest(t *testing.T) {
	tests := []struct {
		name        string
		preferences []int
		k           int
		expected    int
	}{
		{
			name:        "Example 1",
			preferences: []int{3, 2, 1, 5, 6, 4},
			k:           2,
			expected:    5,
		},
		{
			name:        "Example 2",
			preferences: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:           4,
			expected:    4,
		},
		{
			name:        "Single element",
			preferences: []int{42},
			k:           1,
			expected:    42,
		},
		{
			name:        "Duplicate values",
			preferences: []int{5, 5, 5, 5, 5},
			k:           3,
			expected:    5,
		},
		{
			name:        "Negative values",
			preferences: []int{-1, -5, -3, -2, -4},
			k:           2,
			expected:    -2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetKthLargest(test.preferences, test.k)
			if result != test.expected {
				t.Errorf("GetKthLargest(%v, %d) = %d, expected %d", test.preferences, test.k, result, test.expected)
			}
		})
	}
}
