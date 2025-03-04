package calc

func Eval(num1, num2 float64, op string) float64 {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Ошибка: деление на ноль невозможно.")
		} else {
			return num1 / num2
		}
	default:
		panic("Ошибка: неизвестная операция.")
	}
}
