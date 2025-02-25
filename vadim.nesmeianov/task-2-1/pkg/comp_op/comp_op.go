package comp_op

import (
	"errors"
	"fmt"
)

type Operand int

const (
	BiggerOrEqual Operand = iota
	LessOrEqual
	Bigger
	Less
	OperandFailed
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Compare[T Ordered](op Operand, a1, a2 T) (bool, error) {
	switch op {
	case BiggerOrEqual:
		return a1 >= a2, nil
	case LessOrEqual:
		return a1 <= a2, nil
	case Bigger:
		return a1 > a2, nil
	case Less:
		return a1 < a2, nil
	default:
		return false, errors.New("unknown operand")
	}
}

func OperandFromString(str string) (Operand, error) {
	switch str {
	case ">=":
		return BiggerOrEqual, nil

	case "<=":
		return LessOrEqual, nil

	case ">":
		return Bigger, nil

	case "<":
		return Less, nil

	default:
		errStr := fmt.Sprintf("failed to parse operand %s", str)
		return OperandFailed, errors.New(errStr)
	}
}
