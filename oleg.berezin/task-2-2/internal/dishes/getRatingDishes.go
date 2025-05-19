package dishes

import (
	"errors"
	"fmt"
)

func getRatingDish() (int, error) {
	var ratingDish int

	_, err := fmt.Scan(&ratingDish)
	if err != nil {
		return -1, err
	}

	if ratingDish < -10000 || ratingDish > 10000 {
		return -1, errors.New("wrong rating dish: it should be between -1000 and 10000")
	}

	return ratingDish, nil
}

func GetRatingDishes(numDishes int) ([]int, error) {
	var ratingDishes []int

	for range numDishes {
		rating, err := getRatingDish()
		if err != nil {
			return ratingDishes, err
		}

		ratingDishes = append(ratingDishes, rating)
	}

	return ratingDishes, nil
}
