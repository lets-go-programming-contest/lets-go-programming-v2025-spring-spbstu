package main

import (
	"fmt"

	input "github.com/yanelox/task-1/internal/Input"
	"github.com/yanelox/task-1/internal/calc"
)

func main() {
	var Num1, Num2, Res float64
	var Op string

	fmt.Println("Enter first number: ")
	Num1 = input.ScanFloatNum()

	fmt.Println("Enter operation (+, -, *, /): ")
	Op = input.ScanOperation()

	fmt.Println("Enter second number: ")
	Num2 = input.ScanFloatNum()

	Res, err := calc.Eval(Num1, Num2, Op)

	if err != nil {
		fmt.Println("Calcualtion error:", err.Error())
	} else {
		ResStr := fmt.Sprintf("%f %s %f = %f", Num1, Op, Num2, Res)
		fmt.Println(ResStr)
	}
}
