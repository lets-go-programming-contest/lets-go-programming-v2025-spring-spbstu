package input

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Operation int

const (
	GreaterOrEqual Operation = iota
	LessOrEqual
)

type TemperatureRequest struct {
	Op   Operation
	Temp int
}

type Department struct {
	EmployeeCount int
	Requests      []TemperatureRequest
}

type InputProcessor struct {
	scanner *bufio.Scanner
}

func NewInputProcessor(scanner *bufio.Scanner) *InputProcessor {
	return &InputProcessor{scanner: scanner}
}

func (ip *InputProcessor) ReadDepartments() ([]Department, error) {
	n, err := ip.readInt()
	if err != nil {
		return nil, err
	}

	departments := make([]Department, 0, n)

	for i := 0; i < n; i++ {
		k, err := ip.readInt()
		if err != nil {
			return nil, err
		}

		requests := make([]TemperatureRequest, 0, k)
		for j := 0; j < k; j++ {
			req, err := ip.readTemperatureRequest()
			if err != nil {
				return nil, err
			}
			requests = append(requests, req)
		}

		departments = append(departments, Department{
			EmployeeCount: k,
			Requests:      requests,
		})
	}

	return departments, nil
}

func (ip *InputProcessor) readInt() (int, error) {
	if !ip.scanner.Scan() {
		return 0, errors.New("failed to read input")
	}
	return strconv.Atoi(ip.scanner.Text())
}

func (ip *InputProcessor) readTemperatureRequest() (TemperatureRequest, error) {
	if !ip.scanner.Scan() {
		return TemperatureRequest{}, errors.New("failed to read input")
	}
	return ParseTemperatureRequest(ip.scanner.Text())
}

func ParseTemperatureRequest(line string) (TemperatureRequest, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return TemperatureRequest{}, errors.New("invalid input format")
	}

	var op Operation
	switch parts[0] {
	case ">=":
		op = GreaterOrEqual
	case "<=":
		op = LessOrEqual
	default:
		return TemperatureRequest{}, errors.New("invalid operation")
	}

	temp, err := strconv.Atoi(parts[1])
	if err != nil {
		return TemperatureRequest{}, err
	}

	return TemperatureRequest{Op: op, Temp: temp}, nil
}
