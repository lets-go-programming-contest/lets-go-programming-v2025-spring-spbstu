package choiseDishes

import (
	"errors"
	"fmt"

	"choise/internal/input"
	"choise/internal/heap"
)

func Run() error {
	fmt.Print("Enter the number of dishes (1-10000): ");
	N, err := input.InputNumber()

	if err != nil {
        return err
    }

	if !input.CheckRange(N, 1, 10000) {
		return errors.New("Number of dishes out of range\n")
	}

	preferences, err := input.InputPreferences(N)
	if err != nil {
		return err
	}

	fmt.Print("Enter the value of k: ")
	k, err := input.InputNumber()

	if err != nil {
        return err
    }

	if !input.CheckRange(k, 1, N) {
		return errors.New("Number of dishes out of range\n")
	}

	result := heap.FindKthLargest(preferences, k)
	fmt.Printf("The %d-th most preferred dish has a score of: %d\n", k, result)

	return nil

}