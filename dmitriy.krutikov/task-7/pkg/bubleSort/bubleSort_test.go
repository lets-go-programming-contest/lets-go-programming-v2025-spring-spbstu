package bubbleSort

import (
	"reflect"
	"testing"
)

type testCase struct {
	name     string
	input    []int
	expected []int
}

var testTable = []testCase{ 
	{ 
		name:     "empty slice",
		input:    []int{},
		expected: []int{},
	}, 
	{ 
		name:     "already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	}, 
	{ 
		name:     "reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "random order",
		input:    []int{3, 1, 4, 5, 2},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "with duplicates",
		input:    []int{3, 2, 1, 2, 3},
		expected: []int{1, 2, 2, 3, 3},
	},
} 


func TestBubbleSort(t *testing.T) {
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			bubbleSort(inputCopy)

			if !reflect.DeepEqual(inputCopy, tc.expected) {
				t.Errorf("For test case '%s':\nInput:    %v\nExpected: %v\nGot:      %v",
					tc.name, tc.input, tc.expected, inputCopy)
			}
		})
	}
}