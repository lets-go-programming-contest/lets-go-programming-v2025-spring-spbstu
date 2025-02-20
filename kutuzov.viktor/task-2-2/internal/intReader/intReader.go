package intReader

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func Read() (int, error) {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		err = errors.New("error: input value is not int")
		return 0, err
	}

	return int(number), nil
}
