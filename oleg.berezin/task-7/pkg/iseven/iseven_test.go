package iseven

import "testing"

func TestIsEven(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{21, false},
		{37, false},
		{0, true},
		{1, false},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IsEven(tc.input)
			if result != tc.expected {
				t.Errorf("%d %% 2: expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
