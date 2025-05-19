package dishes

import (
	"errors"
	"fmt"
)

func GetTodayDish(numDishes int) (int, error) {
	var ratingDish int

	_, err := fmt.Scan(&ratingDish)
	if err != nil {
		return -1, err
	}

	if ratingDish < 1 || ratingDish > numDishes {
		return -1, errors.New("wrong rating dish: it should be between 1 and dishes num")
	}

	return ratingDish, nil
}
