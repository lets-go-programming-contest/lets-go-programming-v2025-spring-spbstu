package main

import (
	"errors"
	"fmt"
	"os"
)

type Op string

const (
	LE Op = "<="
	GE Op = ">="
)

func readOperation() (Op, error) {
	var opStr string
	_, err := fmt.Scan(&opStr)
	if err != nil {
		return "", err
	}
	switch opStr {
	case ">=":
		return GE, nil
	case "<=":
		return LE, nil
	default:
		return "", errors.New("unknown operator \"" + opStr + "\"")
	}
}

func readInt() (int, error) {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, errors.New("error in scan " + err.Error())
	}
	return num, nil
}

type Constriction struct {
	op  Op
	num int
}

func readConstriction() (Constriction, error) {
	op, err := readOperation()
	if err != nil {
		return Constriction{}, err
	}
	num, err := readInt()
	if err != nil {
		return Constriction{}, err
	}
	return Constriction{op, num}, nil
}

type Range struct {
	lowerBound int
	upperBound int
}

func updateRange(ran Range, con Constriction) Range {
	switch con.op {
	case LE:
		if con.num < ran.lowerBound {
			return Range{-1, -1}
		} else if con.num < ran.upperBound {
			return Range{ran.lowerBound, con.num}
		}
		return ran

	case GE:
		if con.num > ran.upperBound {
			return Range{-1, -1}
		} else if con.num > ran.lowerBound {
			return Range{con.num, ran.upperBound}
		}
		return ran

	default:
		return Range{-1, -1}
	}
}

func process() error {
	N, err := readInt()
	if err != nil {
		return err
	}
	for n := 0; n != N; n++ {
		K, err := readInt()
		if err != nil {
			return err
		}
		ran := Range{15, 30}
		for k := 0; k != K; k++ {
			con, err := readConstriction()
			if err != nil {
				return err
			}
			ran = updateRange(ran, con)
			fmt.Printf("%v\n", ran.lowerBound)
		}
	}
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
	}
}
