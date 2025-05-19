package main

import (
	"fmt"

	"task-2-2/internal/dishes"
	"task-2-2/internal/intheap"
)

func main() {
	numDishes, err := dishes.GetNumDish()
	if err != nil {
		fmt.Printf("error during GetNumDish\n")
		panic(err)
	}

	ratingDishes, err := dishes.GetRatingDishes(numDishes)
	if err != nil {
		fmt.Printf("error during GetRatingDishes\n")
		panic(err)
	}

	todayDish, err := dishes.GetTodayDish(numDishes)
	if err != nil {
		fmt.Printf("error during GetTodayDish\n")
		panic(err)
	}

	result := intheap.GetK(ratingDishes, todayDish)
	fmt.Println(result)
}
