package preference

import (
	"testing"
)

func TestGetKthPrefered(t *testing.T) {
	tests := []struct {
		name        string
		preferences []int
		k           int
		expected    int
	}{
		{
			name:        "Test 1",
			preferences: []int{3, 2, 1, 5, 6, 4},
			k:           2,
			expected:    5,
		},
		{
			name:        "Test 2",
			preferences: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:           4,
			expected:    4,
		},
		{
			name:        "Equal preferences",
			preferences: []int{5, 5, 5, 5, 5},
			k:           3,
			expected:    5,
		},
		{
			name:        "Single",
			preferences: []int{42},
			k:           1,
			expected:    42,
		},
		{
			name:        "Negative",
			preferences: []int{-1, -5, -3, -2, -4},
			k:           2,
			expected:    -2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetKthPrefered(test.preferences, test.k)
			if result != test.expected {
				t.Errorf("GetKthPrefered(%v, %d) = %d, expected %d", test.preferences, test.k, result, test.expected)
			}
		})
	}
}
