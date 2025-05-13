package input

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedData InputData
		expectError  bool
	}{
		{
			name:         "Valid input - Example 1",
			input:        "6\n3 2 1 5 6 4\n2\n",
			expectedData: InputData{Preferences: []int{3, 2, 1, 5, 6, 4}, K: 2},
			expectError:  false,
		},
		{
			name:         "Valid input - Example 2",
			input:        "9\n3 2 3 1 2 4 5 5 6\n4\n",
			expectedData: InputData{Preferences: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, K: 4},
			expectError:  false,
		},
		{
			name:        "Invalid: Too many dishes",
			input:       "10001\n1 2 3\n1\n",
			expectError: true,
		},
		{
			name:        "Invalid: too high",
			input:       "3\n3 10001 1\n1\n",
			expectError: true,
		},
		{
			name:        "Invalid: too low",
			input:       "3\n3 -10001 1\n1\n",
			expectError: true,
		},
		{
			name:        "Invalid: K out of range (too high)",
			input:       "3\n3 2 1\n4\n",
			expectError: true,
		},
		{
			name:        "Invalid: K out of range (too low)",
			input:       "3\n3 2 1\n0\n",
			expectError: true,
		},
		{
			name:        "Invalid: Mismatch in number of dishes",
			input:       "5\n1 2 3\n1\n",
			expectError: true,
		},
		{
			name:        "Invalid: Non-numeric input",
			input:       "3\n1 x 3\n1\n",
			expectError: true,
		},
		{
			name:        "Invalid: Incomplete input",
			input:       "3\n1 2 3\n",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.input)
			data, err := ReadInput(reader)

			if test.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(data, test.expectedData) {
					t.Errorf("Expected data %+v, got %+v", test.expectedData, data)
				}
			}
		})
	}
}
