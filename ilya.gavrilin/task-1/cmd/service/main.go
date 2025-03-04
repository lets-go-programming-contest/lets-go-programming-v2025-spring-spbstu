package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	for {
		num1, err := readNumber("Enter first value: ")
		if err != nil {
			fmt.Println(err)
			continue
		}

		op, err := readOperator("Select operation (+, -, *, /): ")
		if err != nil {
			fmt.Println(err)
			continue
		}

		num2, err := readNumber("Enter second value: ")
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(num1, num2, op)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Result: %g %s %g = %g\n", num1, op, num2, result)

		var again string
		fmt.Print("Continue calculation? (y/n): ")
		if _, err := fmt.Scan(&again); err != nil {
			again = "n"
		}
		if again != "y" && again != "Y" {
			break
		}
	}
}

func readNumber(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	if _, err := fmt.Scan(&input); err != nil {
		return 0, errors.New("Incorrect input. Try again.")
	}
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("Incorrect value. Please, type numeric one.")
	}
	return num, nil
}

func readOperator(prompt string) (string, error) {
	var op string
	fmt.Print(prompt)
	if _, err := fmt.Scan(&op); err != nil {
		return "", errors.New("Incorrect input. Try again.")
	}
	if op != "+" && op != "-" && op != "*" && op != "/" {
		return "", errors.New("Incorrect operation. Please, use: +, -, * or /.")
	}
	return op, nil
}

func calculate(a, b float64, op string) (float64, error) {
	if op == "/" && b == 0 {
		return 0, errors.New("Error: division by zero.")
	}

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	}
	return 0, errors.New("Unknown error.")
}
