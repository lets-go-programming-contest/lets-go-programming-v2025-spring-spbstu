package main

import (
	"fmt"
	"strconv"
)

func scanOperand() float64 {
	for {
		fmt.Print("Enter a float: ")

		var input string
		fmt.Scanln(&input)

		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num
		}
		fmt.Print("Invalid input. Please enter a valid float.")
	}
}

func scanOperator() string {
	for {
		var operator string
		fmt.Scan(&operator)
		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			return operator
		}

		fmt.Print("Error: please enter one of the allowed operators ('+', '-', '*', '/'): ")
	}
}

func main() {
	fmt.Print("Enter the first operand: ")
	operand1 := scanOperand()

	fmt.Print("Enter the second operand: ")
	operand2 := scanOperand()

	fmt.Print("Enter the operator ('+', '-', '*', or '/'): ")
	operator := scanOperator()

	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Error: division by zero is not allowed.")
			return
		}
		result = operand1 / operand2
	}

	fmt.Printf("Result: %.2f\n", result)
}
