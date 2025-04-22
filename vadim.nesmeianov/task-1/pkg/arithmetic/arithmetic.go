package arithmetic

import (
	"errors"
	"fmt"
	"math"
)

const Epsilon = 10e-12

type OperandMethod func(float64, float64) (float64, error)

type Operand interface {
	Perform(float64, float64) (float64, error)
	GetOperandString() string
}

type Op struct {
	calc  OperandMethod
	opStr string
}

func (operand Op) Perform(a1, a2 float64) (float64, error) {
	return operand.calc(a1, a2)
}

func (operand Op) GetOperandString() string {
	return operand.opStr
}

func StringToOperand(op string) (Op, error) {
	var fun OperandMethod

	switch op {
	case "+":
		fun = plus
	case "-":
		fun = minus
	case "*":
		fun = mult
	case "/":
		fun = div
	default:
		errorStr := fmt.Sprintf("unknown operator: %s", op)
		return Op{}, errors.New(errorStr)
	}

	ret := Op{fun, op}

	return ret, nil
}

func plus(a, b float64) (float64, error) {
	return a + b, nil
}

func minus(a, b float64) (float64, error) {
	return a - b, nil
}

func mult(a, b float64) (float64, error) {
	return a * b, nil
}

func div(a, b float64) (float64, error) {
	if isEqual(b, 0) {
		sign := getSign(a)
		return math.Inf(sign), errors.New("divider is zero")
	} else {
		return a / b, nil
	}
}

func isEqual(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}

func getSign(a float64) int {
	if a > 0 {
		return 1
	} else {
		return -1
	}
}
