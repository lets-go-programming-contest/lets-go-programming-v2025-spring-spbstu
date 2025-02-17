package input

import (
	"fmt"
	"strings"
)

func ScanFloatNum() float64 {
	var Num float64

	_, err := fmt.Scan(&Num)

	for err != nil {
		var garbage string
		fmt.Scanln(&garbage)

		fmt.Println("Incorrect input, should be floating point number, try again")

		_, err = fmt.Scan(&Num)
	}

	return Num
}

func ScanOperation() string {
	var Op string
	var AvailableOperations = "+-*/"

	_, err := fmt.Scan(&Op)

	for (err != nil) || (!strings.Contains(AvailableOperations, Op)) {
		fmt.Println("Incorrect input, should be '+', '-', '*' or '/', try again")

		_, err = fmt.Scan(&Op)
	}

	return Op
}
