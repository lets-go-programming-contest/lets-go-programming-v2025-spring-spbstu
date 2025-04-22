package main

import (
	"fmt"

	input "github.com/yanelox/task-1/internal/Input"
	"github.com/yanelox/task-1/internal/calc"
)

func main() {
	fmt.Println("Enter first number: ")
	num1 := input.ScanFloatNum()

	fmt.Println("Enter operation (+, -, *, /): ")
	op := input.ScanOperation()

	fmt.Println("Enter second number: ")
	num2 := input.ScanFloatNum()

	res, err := calc.Eval(num1, num2, op)

	if err != nil {
		fmt.Println("Calcualtion error:", err.Error())
	} else {
		fmt.Printf("%f %s %f = %f\n", num1, op, num2, res)
	}
}
