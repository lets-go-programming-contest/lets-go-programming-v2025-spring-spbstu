package uintReader

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func Read() (uint64, error) {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		err = errors.New("error: input value is not uint")
		return 0, err
	}

	return number, nil
}
