package temperature

import (
	"errors"
	"fmt"
)

func getSign() (bool, error) {
	var s string

	if _, err := fmt.Scan(&s); err != nil {
		return false, err
	}

	switch s {
	case "<=":
		return true, nil
	case ">=":
		return false, nil
	default:
		return false, errors.New("invalid sign: must be \"<=\" or \">=\"")
	}
}

func getT() (int, error) {
	var t int

	if _, err := fmt.Scan(&t); err != nil {
		return 0, err
	}

	if t < 15 || t > 30 {
		err := errors.New("wrong temperatyre number: it should be between 15 and 30")
		return -1, err
	}

	return t, nil
}

func GetTemperature() (Temperature, error) {
	less, err := getSign()
	if err != nil {
		return Temperature{}, err
	}

	t, err := getT()
	if err != nil {
		return Temperature{}, err
	}

	return Temperature{
		T:    t,
		Less: less,
	}, nil
}
