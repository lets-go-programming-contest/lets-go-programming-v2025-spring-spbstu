package calc

import "fmt"

func Eval[T int | float32 | float64](num1 T, num2 T, op string) (T, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		} else {
			return num1 / num2, nil
		}
	default:
		return 0, fmt.Errorf("unknown operation")
	}
}
