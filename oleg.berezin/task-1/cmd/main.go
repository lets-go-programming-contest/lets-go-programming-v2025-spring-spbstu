package main

import (
	"calc/pkg/calc"
	"calc/pkg/io"
	"fmt"
)

func main() {
	fmt.Print("Введите первое число: ")
	num1 := io.GetNum()

	fmt.Print("Выберите операцию (+, -, *, /): ")
	op := io.GetOp()

	fmt.Print("Введите второе число: ")
	num2 := io.GetNum()

	fmt.Printf("Результат: %g %s %g = %g\n", num1, op, num2, calc.Eval(num1, num2, op))
}
