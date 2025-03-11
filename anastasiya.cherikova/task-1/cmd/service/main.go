package main

import (
	"bufio"   // Импортируем пакет bufio для работы с буферизованным вводом
	"fmt"     // Импортируем пакет fmt для форматированного ввода/вывода
	"os"      // Импортируем пакет os для работы с операционной системой
	"strings" // Импортируем пакет strings для работы со строками
)

// Функция readNumber запрашивает у пользователя ввод числа
func readNumber(prompt string) float64 {
	scanner := bufio.NewScanner(os.Stdin) // Создаем новый сканер для считывания ввода из стандартного ввода
	for {
		fmt.Print(prompt)
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text()) // Удаляем лишние пробелы из введенной строки

		var num float64
		_, err := fmt.Sscanf(text, "%f", &num) // Пытаемся считать число из текста
		if err == nil {
			return num
		}
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
	}
}

// Функция readOperation запрашивает у пользователя выбор операции (+, -, *, /)
func readOperation() string {
	scanner := bufio.NewScanner(os.Stdin) // Создаем новый сканер для считывания ввода из стандартного ввода
	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		scanner.Scan()
		op := strings.TrimSpace(scanner.Text()) // Удаляем лишние пробелы из введенной строки

		if len(op) == 1 && strings.ContainsAny(op, "+-*/") {
			return op
		}
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}

func main() {
	num1 := readNumber("Введите первое число: ")
	op := readOperation()

	var num2 float64
	if op == "/" {
		for {
			num2 = readNumber("Введите второе число: ")
			if num2 != 0 {
				break
			}
			fmt.Println("Ошибка: деление на ноль невозможно.")
		}
	} else {
		num2 = readNumber("Введите второе число: ")
	}

	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}

	fmt.Printf("Результат: %v %v %v = %v\n", num1, op, num2, result)
}
