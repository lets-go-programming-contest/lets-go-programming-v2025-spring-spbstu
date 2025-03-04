package io

import (
	"fmt"
	"strings"
)

func GetNum() float64 {
	var num float64

	_, err := fmt.Scanln(&num)

	for err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение формата float64.")

		_, err = fmt.Scanln(&num)
	}

	return num
}

func GetOp() string {
	const operations = "+-*/"

	var op string

	_, err := fmt.Scanln(&op)

	for (err != nil) || (!strings.Contains(operations, op)) {
		fmt.Println("Пожалуйста, используйте символы +, -, * или /.")

		_, err = fmt.Scanln(&op)
	}

	return op
}
