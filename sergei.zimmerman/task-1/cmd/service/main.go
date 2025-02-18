package main

import (
	"errors"
	"fmt"
	"os"
)

type Op int

const (
	Add Op = iota
	Sub
	Mul
	Div
)

func parseOp(str string) (Op, error) {
	switch str {
	case "+":
		return Add, nil
	case "-":
		return Sub, nil
	case "*":
		return Mul, nil
	case "/":
		return Div, nil
	}

	return Add, errors.New("invalid op, only '+', '-', '*' or '/' are valid options")
}

func apply(a int, b int, f Op) (int, error) {
	switch f {
	case Add:
		return a + b, nil
	case Sub:
		return a - b, nil
	case Mul:
		return a * b, nil
	case Div:
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	}

	panic("unreachable")
}

func readNum() (int, error) {
	var a int
	_, err := fmt.Scan(&a)
	return a, err
}

func run() error {
	fmt.Print("enter first value: ")
	a, err := readNum()

	if err != nil {
		return err
	}

	fmt.Print("enter second value: ")
	b, err := readNum()
	if err != nil {
		return err
	}

	fmt.Print("select op (+, -, *, /): ")
	var opstr string
	_, err = fmt.Scanf("%s", &opstr)
	if err != nil {
		return err
	}

	f, err := parseOp(opstr)
	if err != nil {
		return err
	}

	result, err := apply(a, b, f)
	if err != nil {
		return err
	}

	fmt.Printf("result: %d %s %d = %d\n", a, opstr, b, result)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
