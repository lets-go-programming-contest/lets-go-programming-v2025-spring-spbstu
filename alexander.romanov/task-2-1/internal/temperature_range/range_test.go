package temperature_range

import (
	"task-2-1/internal/input"
	"testing"
)

func TestOfficeTemperatures(t *testing.T) {
	tests := []struct {
		name    string
		updates []struct {
			op   input.Operation
			temp int
		}
		expected *int
	}{
		{
			name: "Initial range",
			updates: []struct {
				op   input.Operation
				temp int
			}{},
			expected: intPtr(MinTemperature),
		},
		{
			name: "Single >= update",
			updates: []struct {
				op   input.Operation
				temp int
			}{
				{input.GreaterOrEqual, 20},
			},
			expected: intPtr(20),
		},
		{
			name: "Single <= update",
			updates: []struct {
				op   input.Operation
				temp int
			}{
				{input.LessOrEqual, 25},
			},
			expected: intPtr(MinTemperature),
		},
		{
			name: "Multiple valid updates",
			updates: []struct {
				op   input.Operation
				temp int
			}{
				{input.GreaterOrEqual, 18},
				{input.LessOrEqual, 23},
				{input.GreaterOrEqual, 20},
				{input.LessOrEqual, 27},
				{input.LessOrEqual, 21},
			},
			expected: intPtr(20),
		},
		{
			name: "Invalid: Conflicting updates",
			updates: []struct {
				op   input.Operation
				temp int
			}{
				{input.GreaterOrEqual, 18},
				{input.LessOrEqual, 23},
				{input.GreaterOrEqual, 20},
				{input.LessOrEqual, 27},
				{input.LessOrEqual, 21},
				{input.GreaterOrEqual, 28},
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTemperatureRange()
			result := tr.GetOptimal()
			for _, update := range tt.updates {
				if !tr.Update(update.op, update.temp) {
					result = nil
					break
				}
				result = tr.GetOptimal()
			}

			assertOptimalTemp(t, tt.expected, result)
		})
	}
}

func assertOptimalTemp(t *testing.T, expected, actual *int) {
	t.Helper()
	if expected == nil && actual != nil {
		t.Error("Expected nil, got non-nil")
	} else if expected != nil && actual == nil {
		t.Error("Expected non-nil, got nil")
	} else if expected != nil && actual != nil && *expected != *actual {
		t.Errorf("Expected %d, got %d", *expected, *actual)
	}
}

func intPtr(i int) *int {
	return &i
}
