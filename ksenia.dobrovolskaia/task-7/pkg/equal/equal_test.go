package equal

import "testing"

func TestIsEqual(t *testing.T) {
	testCases := []struct {
		inputA   int
		inputB   int
		expected bool
	}{
		{1, 2, false},
		{-1, -2, false},
		{2, 1, false},
		{1, 1, true},
		{0, 0, true},
		{-1, -1, true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IsEqual(tc.inputA, tc.inputB)
			if result != tc.expected {
				t.Errorf("%d == %d: expected %v, but got %v", tc.inputA, tc.inputB, tc.expected, result)
			}
		})
	}
}
