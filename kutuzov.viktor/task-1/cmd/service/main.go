package main

import (
	"fmt"
)

func main() {

	var num1 float64
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	var operation string
	fmt.Print("Выберите операцию (+, -, *, /): ")
	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Ошибка ввода операции.")
		return
	}

	var num2 float64
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	switch operation {
	case "+":
		fmt.Println("Результат:", num1, operation, num2, "=", num1+num2)

	case "-":
		fmt.Println("Результат:", num1, operation, num2, "=", num1-num2)

	case "*":
		fmt.Println("Результат:", num1, operation, num2, "=", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			fmt.Println("Результат:", num1, operation, num2, "=", num1/num2)
		}

	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}
