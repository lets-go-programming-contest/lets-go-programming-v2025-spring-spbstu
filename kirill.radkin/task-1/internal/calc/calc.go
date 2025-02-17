package calc

import "fmt"

func Eval[T int | float32 | float64](Num1 T, Num2 T, Op string) (T, error) {
	var Res T = 0
	var err error

	switch Op {
	case "+":
		Res = Num1 + Num2
	case "-":
		Res = Num1 - Num2
	case "*":
		Res = Num1 * Num2
	case "/":
		if Num2 == 0 {
			err = fmt.Errorf("division by zero")
		} else {
			Res = Num1 / Num2
		}
	default:
		err = fmt.Errorf("unknown operation")
	}

	return Res, err
}
