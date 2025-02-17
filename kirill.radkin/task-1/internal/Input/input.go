package input

import (
	"fmt"
	"strings"
)

func ScanFloatNum() float64 {
	var num float64

	_, err := fmt.Scan(&num)

	for err != nil {
		var garbage string
		fmt.Scanln(&garbage)

		fmt.Println("Incorrect input, should be floating point number, try again")

		_, err = fmt.Scan(&num)
	}

	return num
}

func ScanOperation() string {
	var op string
	var AvailableOperations = "+-*/"

	_, err := fmt.Scan(&op)

	for (err != nil) || (!strings.Contains(AvailableOperations, op)) {
		fmt.Println("Incorrect input, should be '+', '-', '*' or '/', try again")

		_, err = fmt.Scan(&op)
	}

	return op
}
