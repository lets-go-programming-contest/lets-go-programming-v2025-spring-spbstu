package calculate

import (
	"errors"
)

func Calculate(val1, val2 float64, op string) (float64, error) {

	var ans float64
	var err error
	switch op {
	case "+":
		ans = val1 + val2

	case "-":
		ans = val1 - val2

	case "*":
		ans = val1 * val2

	case "/":
		if val2 == 0 {
			err = errors.New("error: division by zero is forbidden")
		} else {
			ans = val1 / val2
		}

	default:
		err = errors.New("error: bad op argument")
	}

	return ans, err
}
