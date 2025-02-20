package readFloat

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Read() (float64, error) {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Insert num:  ")

	input, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		err = errors.New("error: input value is not float")
		return 0, err
	}

	return number, nil
}
