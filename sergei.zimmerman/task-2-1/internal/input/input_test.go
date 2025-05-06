package input

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseTemperatureRequest(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedOp   Operation
		expectedTemp int
		expectErr    bool
	}{
		{"Valid >=", ">= 20", GreaterOrEqual, 20, false},
		{"Valid <=", "<= 25", LessOrEqual, 25, false},
		{"Invalid op", "= 20", 0, 0, true},
		{"Invalid temp", ">= abc", 0, 0, true},
		{"Empty", "", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := ParseTemperatureRequest(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if req.Op != tt.expectedOp {
				t.Errorf("Expected op %v, got %v", tt.expectedOp, req.Op)
			}
			if req.Temp != tt.expectedTemp {
				t.Errorf("Expected temp %d, got %d", tt.expectedTemp, req.Temp)
			}
		})
	}
}

func TestReadDepartments(t *testing.T) {
	inputStr := `2
1
>= 30
6
>= 18
<= 23
>= 20
<= 27
<= 21
>= 28`

	scanner := bufio.NewScanner(strings.NewReader(inputStr))
	processor := NewInputProcessor(scanner)

	departments, err := processor.ReadDepartments()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(departments) != 2 {
		t.Fatalf("Expected 2 departments, got %d", len(departments))
	}

	if departments[0].EmployeeCount != 1 {
		t.Errorf("Expected 1 employee in first department, got %d", departments[0].EmployeeCount)
	}
	if len(departments[0].Requests) != 1 {
		t.Fatalf("Expected 1 request in first department, got %d", len(departments[0].Requests))
	}
	if departments[0].Requests[0].Op != GreaterOrEqual || departments[0].Requests[0].Temp != 30 {
		t.Errorf("First request mismatch: got %v %d", departments[0].Requests[0].Op, departments[0].Requests[0].Temp)
	}

	if departments[1].EmployeeCount != 6 {
		t.Errorf("Expected 6 employees in second department, got %d", departments[1].EmployeeCount)
	}
	if len(departments[1].Requests) != 6 {
		t.Fatalf("Expected 6 requests in second department, got %d", len(departments[1].Requests))
	}
}

func TestReadInt(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader("123\nabc"))
	processor := NewInputProcessor(scanner)

	val, err := processor.readInt()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if val != 123 {
		t.Errorf("Expected 123, got %d", val)
	}

	_, err = processor.readInt()
	if err == nil {
		t.Error("Expected error for non-integer input, got nil")
	}
}
