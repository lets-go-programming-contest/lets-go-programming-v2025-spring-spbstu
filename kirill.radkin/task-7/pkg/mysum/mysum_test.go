package mysum_test

import (
	"testing"

	"github.com/yanelox/task-7/pkg/mysum"
)

type TestCase struct {
	name     string
	a, b     int
	expected int
}

func TestMySum(t *testing.T) {
	tests := []TestCase {
		{"positive numbers", 1, 2, 3},
		{"zero", 0, 0, 0},
		{"negative numbers", -1, -2, -3},
		{"mixed", -1, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mysum.MySum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("MySum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}