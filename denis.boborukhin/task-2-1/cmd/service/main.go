package main

import (
	"fmt"
)

func scanOperator() (string, error) {
	var op string
	_, err := fmt.Scan(&op)
	if err != nil {
		return "", fmt.Errorf("failed to read operator: %w", err)
	}
	if op != "<=" && op != ">=" {
		return "", fmt.Errorf("invalid operator: %s", op)
	}
	return op, nil
}

func scanNumber() (int, error) {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, fmt.Errorf("failed to read number: %w", err)
	}
	return num, nil
}

type Condition struct {
	operator string
	value    int
}

func scanCondition() (Condition, error) {
	op, err := scanOperator()
	if err != nil {
		return Condition{}, err
	}
	num, err := scanNumber()
	if err != nil {
		return Condition{}, err
	}
	return Condition{op, num}, nil
}

type TemperatureRange struct {
	min int
	max int
}

func adjustRange(current TemperatureRange, cond Condition) TemperatureRange {
	switch cond.operator {
	case "<=":
		if cond.value < current.min {
			return TemperatureRange{-1, -1}
		}
		if cond.value < current.max {
			return TemperatureRange{current.min, cond.value}
		}
		return current

	case ">=":
		if cond.value > current.max {
			return TemperatureRange{-1, -1}
		}
		if cond.value > current.min {
			return TemperatureRange{cond.value, current.max}
		}
		return current

	default:
		return TemperatureRange{-1, -1}
	}
}

func execute() error {
	departmentsCount, err := scanNumber()
	if err != nil {
		return fmt.Errorf("failed to read departments count: %w", err)
	}

	for i := 0; i < departmentsCount; i++ {
		employeesCount, err := scanNumber()
		if err != nil {
			return fmt.Errorf("failed to read employees count: %w", err)
		}

		const (
			MinTemp = 15
			MaxTemp = 30
		)
		rangeTemp := TemperatureRange{MinTemp, MaxTemp}
		for j := 0; j < employeesCount; j++ {
			cond, err := scanCondition()
			if err != nil {
				return fmt.Errorf("failed to read condition: %w", err)
			}

			rangeTemp = adjustRange(rangeTemp, cond)
			fmt.Println(rangeTemp.min)
		}
	}
	return nil
}

func main() {
	if err := execute(); err != nil {
		fmt.Println("Error occurred:", err)
	}
}
