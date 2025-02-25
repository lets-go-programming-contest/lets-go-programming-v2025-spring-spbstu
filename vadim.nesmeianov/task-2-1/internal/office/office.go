package office

import (
	"errors"
	"fmt"
	"strings"
	"task-2-1/internal/department"
	"task-2-1/pkg/comp_op"
)

type Office struct {
	depNum int
	dep    []department.Department
}

func (instance Office) GetTempInDepartment(id int) int {
	return instance.dep[id].GetTemperature()
}

func (instance *Office) EventEmployeeCame(dep, temp int, op comp_op.Operand) error {
	if dep < instance.depNum {
		return instance.dep[dep].EventEmployeeCame(temp, op)
	}

	errStr := fmt.Sprintf("dep_id = %d is too big", dep)
	return errors.New(errStr)
}

func (instance *Office) SetDepartmentCapasity(num, size int) error {
	if num < instance.depNum {
		instance.dep[num].SetCapacity(size)
		return nil
	}

	errStr := fmt.Sprintf("dep_id = %d is too big", num)
	return errors.New(errStr)
}

func GetOffice(n int) *Office {
	instance := new(Office)
	instance.depNum = n
	instance.dep = make([]department.Department, n)

	for i, _ := range instance.dep {
		instance.dep[i] = department.GetDepartment(i)
	}

	return instance
}

func Run(logOutput bool) error {
	var n int
	err := readValue(&n, "Enter N:", logOutput)
	if err != nil {
		return err
	}

	office := GetOffice(n)

	for i := 0; i < n; i += 1 {
		var k int
		msgStr := fmt.Sprintf("Enter K for department %d", i)
		err = readValue(&k, msgStr, logOutput)
		if err != nil {
			return err
		}

		office.SetDepartmentCapasity(i, k)

		for j := 0; j < k; j += 1 {
			op, temp, err := readRequest()
			if err != nil {
				return err
			}

			err = office.EventEmployeeCame(i, temp, op)
			if err != nil {
				return err
			}

			currTemp := office.GetTempInDepartment(i)
			fmt.Printf("%d\n", currTemp)
		}
	}

	return nil
}

func readRequest() (comp_op.Operand, int, error) {
	var str string
	var k int
	_, err := fmt.Scanf("%s%d\n", &str, &k)
	if err != nil {
		return comp_op.OperandFailed, 0, err
	}

	lexemsArr := strings.Fields(str)
	if len(lexemsArr) != 1 {
		errStr := fmt.Sprintf("the string: \" %s \" has %d lexems, must be 2", str, len(lexemsArr))
		return comp_op.OperandFailed, 0, errors.New(errStr)
	}
	op, err := comp_op.OperandFromString(lexemsArr[0])
	if err != nil {
		return comp_op.OperandFailed, 0, err
	}

	return op, k, nil
}

func readValue[T any](x *T, msg string, logOutput bool) error {
	if logOutput {
		fmt.Println(msg)
	}

	_, err := fmt.Scan(x)
	return err
}
