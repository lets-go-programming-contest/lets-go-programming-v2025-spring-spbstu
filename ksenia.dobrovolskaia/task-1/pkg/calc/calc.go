package calc

import (
	"errors"
)

// Производит арифметическую операцию operation с числами a и b
// Поддерживаемые operation: "+", "-", "*", "/"
func Calculate(a, b int, operation string) (result int, err error) {
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			err = errors.New("Ошибка: деление на ноль невозможно.")
			return
		}
		result = a / b
	default:
		err = errors.New("Некорректная операция.")
	}
	return
}
