package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math"
)

const float64EqualityThreshold = 1e-9

func isFloat64Equals(lhs, rhs float64) bool {
	return math.Abs(lhs-rhs) <= float64EqualityThreshold
}

type CalculatorError struct {
	what string
}

func (e *CalculatorError) Error() string {
	return e.what
}

func parse() (string, float64, float64, error) {
	var err error

	fmt.Printf("\nEnter the first argument:\n")
	var lhs float64
	_, err = fmt.Scan(&lhs)
	if err != nil {
		return "", 0, 0, &CalculatorError{fmt.Sprintf("Cant parse the fisrt argument: %s", err)}
	}

	fmt.Printf("Enter a binary operation:\n")
	var op string
	_, err = fmt.Scan(&op)
	if err != nil {
		return "", 0, 0, &CalculatorError{fmt.Sprintf("Cant parse an oprator: %s", err)}
	}

	fmt.Printf("Enter the second argument:\n")
	var rhs float64
	_, err = fmt.Scan(&rhs)
	if err != nil {
		return "", 0, 0, &CalculatorError{fmt.Sprintf("Cant parse the second argument: %s", err)}
	}

	return op, lhs, rhs, nil
}

func calculate(op string, lhs, rhs float64) (float64, error) {
	switch op {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if isFloat64Equals(rhs, 0.0) {
			return 0, &CalculatorError{fmt.Sprintf("Divizion by zero: %g %s %g", lhs, op, rhs)}
		}
		return lhs / rhs, nil
	default:
		return 0, &CalculatorError{fmt.Sprintf("Unknown operation: %s", op)}
	}
}

func shouldExit() bool {
	c, _, err := keyboard.GetSingleKey()
	if err != nil {
		return false
	}

	if c == 'q' {
		return true
	}

	return false
}

func main() {
	for {
		fmt.Printf("\nNaive calculator application\n" +
			"Currently only 4 operations are supported: +, -, *, /\n" +
			"Press any key to continue, or enter \033[1mq\033[0m to exit:\n")

		if shouldExit() {
			break
		}

		op, lhs, rhs, err := parse()
		if err != nil {
			fmt.Printf("ParseError: %s\n Restarting...\n\n", err)
			continue
		}

		res, err := calculate(op, lhs, rhs)
		if err != nil {
			fmt.Printf("CalculationError: %s\n Restarting...\n\n", err)
			continue
		}

		fmt.Printf("Calculation successed\n"+
			"Result = \033[1m%f\033[0m\n\n", res)
	}
}
