package main

import (
	"fmt"
	"strconv"
)

func readOperand() float64 {
	for {
		var input string
		fmt.Scan(&input)
		num, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return num
		}
		fmt.Print("Некорректное число. Пожалуйста, введите числовое значение: ")
	}
}

func readOperator() string {
	for {
		var input string
		fmt.Scan(&input)
		if input == "+" || input == "-" || input == "*" || input == "/" {
			return input
		}
		fmt.Print("Некорректная операция. Пожалуйста, используйте символы { +, -, *, / }: ")
	}
}

func main() {
	for {
		fmt.Print("Введите первое число: ")
		num1 := readOperand()
		fmt.Print("Выберите арифметическую операцию ( +, -, *, / ): ")
		operator := readOperator()
		fmt.Print("Введите второе число: ")
		num2 := readOperand()

		var result float64
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Ошибка: деление на ноль невозможно.")
				continue
			}
			result = num1 / num2
		}

		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

		var choice string
		fmt.Print("Для продолжения нажмите 'r', для выхода любую другую клавишу: ")
		fmt.Scan(&choice)
		if choice != "r" {
			break
		}
	}
}
