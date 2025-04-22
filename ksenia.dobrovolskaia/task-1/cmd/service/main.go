package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/kseniadobrovolskaia/task-1/internal/calc"
)

func main() {
	a, b, operation, err := readInput()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	result, err := calc.Calculate(a, b, operation)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v %d %v %d = %d\n", color.YellowString("Результат:"), a, operation, b, result)
}

func readInput() (a, b int, operation string, err error) {
	color.Green("Введите первое число: ")
	_, err = fmt.Scan(&a)
	if err != nil {
		err = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	color.Green("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&operation)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		err = errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}
	color.Green("Введите второе число: ")
	fmt.Scan(&b)
	return
}
