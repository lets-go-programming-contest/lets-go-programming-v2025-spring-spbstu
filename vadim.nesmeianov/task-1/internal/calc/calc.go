package calc

import (
	"errors"
	"fmt"
	arithmetic "task-1/pkg/arithmetic"
)

func Run() error {
	a1, a2, err := readValues()
	if err != nil {
		return err
	}

	op, err := readOperator()
	if err != nil {
		return err
	}

	a3, err := op.Perform(a1, a2)
	if err != nil {
		return err
	}

	printAnswer(a1, a2, a3, op)
	return nil
}

func printAnswer(a1, a2, a3 float64, op arithmetic.Operand) {
	fmt.Printf("The expression evaluation:\n%f %s %f = %f\n", a1, op.GetOperandString(), a2, a3)
}

func readOperator() (arithmetic.Operand, error) {
	fmt.Printf("Enter an operator(+, -, /, *):\n")

	var inputStr string
	_, err := fmt.Scan(&inputStr)

	if err != nil {
		errorStr := fmt.Sprintf("unknown operand: %s", inputStr)
		return nil, errors.New(errorStr)
	}

	op, err := arithmetic.StringToOperand(inputStr)

	if err != nil {
		errorStr := fmt.Sprintf("cannot parse operand: %s", inputStr)
		return nil, errors.New(errorStr)
	}

	return op, nil
}

func readValues() (float64, float64, error) {
	fmt.Print("Enter 1st and 2nd value\n")

	var a1, a2 float64
	_, err := fmt.Scan(&a1, &a2)

	return a1, a2, err
}
