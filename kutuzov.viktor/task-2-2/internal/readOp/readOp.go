package readOp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Read() (string, error) {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Insert operator (+, -, *, /): ")

	op, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	op = strings.TrimSpace(op)

	if !(op == "+" || op == "-" || op == "*" || op == "/") {
		err = errors.New("error: incorrect operator")
		return "", err
	}

	return op, nil
}
