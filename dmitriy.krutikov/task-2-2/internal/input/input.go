package input

import (
    "bufio"
    "errors"
    "fmt"
	"os"
	"strconv"
	"strings"
)

func InputNumber() (int, error) {
	var number int
	_, err := fmt.Scanln(&number)

	if err != nil {
		return 0, errors.New("Incorrect number\n")
	}

	return number, nil
}

func CheckRange(number, min, max int) bool  {
	if number >= min && number <= max {
		return true
	}
	return false
}

func InputPreferences(N int) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return nil, errors.New("incorrect input: failed to read the input")
	}

	input = strings.TrimSpace(input)

	prefList := strings.Split(input, " ")

	if len(prefList) != N {
		return nil,  errors.New("unexpected size")
	}

	preferences := make([]int, N, N)
	for i, p := range prefList {
		score, err := strconv.Atoi(p)

		if err != nil {
			return nil, errors.New("invalid input: preference scores must be numbers")
		}

		if !CheckRange(score, -10000, 10000) {
			return nil, errors.New("Number of dishes out of range\n")
		}

		preferences[i] = score
	}

	return preferences, nil
}