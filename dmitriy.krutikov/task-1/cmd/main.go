package main

import (
	"errors"
	"fmt"
)

func isValidOperator(operator string) bool {
	switch operator {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func inputNumber() (float64, error) {
	var number float64
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)

	if err != nil {
		return 0, errors.New("Некорректное число. Пожалуйста, введите числовое значение")
	}

	return number, nil
}

func inputOperator() (string, error) {
	var operator string
	fmt.Print("Введите операцию(+, -, *, /): ")
	_, err := fmt.Scan(&operator)

	if err != nil {
		return "", errors.New("Некорректный ввод. Пожалуйста, введите +, -, *, /")
	}

	if !isValidOperator(operator) {
		return "", errors.New("Некорректный оператор. Пожалуйста, введите +, -, *, /")
	}

	return operator, nil

}

func applyOperation(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("Деление на ноль")			
		}

		return num1 / num2, nil
	default:
		return 0, errors.New("Некорректный оператор. Пожалуйста, введите +, -, *, /")
	}
}

func main() {
	num1, err := inputNumber()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	operator, err := inputOperator()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	num2, err := inputNumber()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	result, err := applyOperation(num1, operator, num2)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Результат: %v %s %v = %v\n", num1, operator, num2, result);
}
