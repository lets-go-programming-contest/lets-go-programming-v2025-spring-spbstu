package dishes

import (
	"errors"
	"fmt"
)

func GetNumDish() (int, error) {
	var numDishes int

	_, err := fmt.Scan(&numDishes)
	if err != nil {
		return -1, err
	}

	if numDishes < 1 || numDishes > 10000 {
		return -1, errors.New("wrong num dishes: it should be between 1 and 10000")
	}

	return numDishes, nil
}
