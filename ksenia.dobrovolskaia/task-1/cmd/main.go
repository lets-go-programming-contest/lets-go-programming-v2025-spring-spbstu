package main

import (
	"errors"
	"fmt"
	"log"
	"task-1/pkg/calc"
)

func main() {
	a, b, operation, err := readInput()
	if err != nil {
		log.Fatal(err)
	}
	result, err := calc.Calculate(a, b, operation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Результат: %d %v %d = %d\n", a, operation, b, result)
}

func readInput() (a, b int, operation string, err error) {
	fmt.Print("Введите первое число: ")
	_, err = fmt.Scan(&a)
	if err != nil {
		err = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&operation)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		err = errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	return
}
