package main

import (
	"errors"
	"fmt"
	"os"
)

type Op string

const (
	ADD Op = "+"
	SUB Op = "-"
	MUL Op = "*"
	DIV Op = "/"
)

func readNumber() (float64, error) {
	var res float64
	_, err := fmt.Scan(&res)
	if err != nil {
		return 0., err
	}
	return res, nil
}

func readOperation() (Op, error) {
	var opStr string
	_, err := fmt.Scan(&opStr)
	if err != nil {
		return "", err
	}
	switch opStr {
	case "+":
		return ADD, nil
	case "-":
		return SUB, nil
	case "*":
		return MUL, nil
	case "/":
		return DIV, nil
	default:
		return "", errors.New("unknown operation: \"" + opStr + "\"")
	}
}

func getResult(lhs float64, rhs float64, op Op) (float64, error) {
	switch op {
	case ADD:
		return lhs + rhs, nil
	case SUB:
		return lhs - rhs, nil
	case MUL:
		return lhs * rhs, nil
	case DIV:
		if rhs == 0 {
			return 0., errors.New("division by zero")
		}
		return lhs / rhs, nil
	default:
		return 0., errors.New("UNDEF operation")
	}
}

func process() error {
	fmt.Printf("enter first number (lhs)\n")
	lhs, err := readNumber()
	if err != nil {
		return err
	}
	fmt.Printf("enter operation (+,-,*,/)\n")
	op, err := readOperation()
	if err != nil {
		return err
	}
	fmt.Printf("enter second number (rhs)\n")
	rhs, err := readNumber()
	if err != nil {
		return err
	}
	res, err := getResult(lhs, rhs, op)
	if err != nil {
		return err
	}
	fmt.Printf("result is %v\n", res)
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
