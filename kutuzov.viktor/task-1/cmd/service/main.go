package main

import (
	"fmt"

	"github.com/kutuzov.viktor/task-1/internal/calculate"
	"github.com/kutuzov.viktor/task-1/internal/print"
	value "github.com/kutuzov.viktor/task-1/internal/readFloat"
	operator "github.com/kutuzov.viktor/task-1/internal/readOp"
)

func main() {
	var val1 float64
	val1, err := value.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	op, err := operator.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	val2, err := value.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	ans, err := calculate.Calculate(val1, val2, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = print.PrintAns(val1, val2, ans, op)
	if err != nil {
		fmt.Println(err)
		return
	}
}
