package department

import (
	"errors"
	"fmt"
	"task-2-1/cmd/ac"
	"task-2-1/cmd/comp_op"
)

type department struct {
	depId         int
	employeesNum  int
	employeesCame int
	ac            ac.Ac
}

type Department interface {
	EventEmployeeCame(temp int, op comp_op.Operand) error
	SetCapacity(int)
	GetTemperature() int
}

func (instance department) GetTemperature() int {
	return instance.ac.GetTemperature()
}

func (instance *department) EventEmployeeCame(temp int, op comp_op.Operand) error {
	if instance.employeesCame >= instance.employeesNum {
		errStr := fmt.Sprintf("department %d is overcrowded", instance.depId)
		return errors.New(errStr)
	}
	return instance.ac.NewRequest(temp, op)
}

func (instance *department) SetCapacity(cap int) {
	instance.employeesNum = cap
}

func GetDepartment(id int) Department {
	instance := new(department)
	instance.depId = id
	instance.employeesCame = 0
	instance.ac = ac.GetAc()
	return instance
}
