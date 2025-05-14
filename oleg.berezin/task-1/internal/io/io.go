package io

import (
	"fmt"
	"strings"
)

func ClearStdin() {
	var remains string
	fmt.Scanln(&remains)
}

func GetNum() float64 {
	var num float64

	_, err := fmt.Scan(&num)

	for err != nil {
		ClearStdin()

		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение формата float64.")

		_, err = fmt.Scan(&num)
	}

	return num
}

func GetOp() string {
	const operations = "+-*/"

	var op string

	_, err := fmt.Scanln(&op)

	for (err != nil) || (!strings.Contains(operations, op)) {
		ClearStdin()

		fmt.Println("Пожалуйста, используйте символы +, -, * или /.")

		_, err = fmt.Scanln(&op)
	}

	return op
}
