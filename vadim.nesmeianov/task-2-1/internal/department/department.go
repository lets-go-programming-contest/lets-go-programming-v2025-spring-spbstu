package department

import (
	"errors"
	"fmt"
	"task-2-1/internal/ac"
	"task-2-1/internal/comp_op"
)

type Department struct {
	depId         int
	employeesNum  int
	employeesCame int
	ac            *ac.Ac
}

func (instance Department) GetTemperature() int {
	return instance.ac.GetTemperature()
}

func (instance *Department) EventEmployeeCame(temp int, op comp_op.Operand) error {
	if instance.employeesCame >= instance.employeesNum {
		errStr := fmt.Sprintf("department %d is overcrowded", instance.depId)
		return errors.New(errStr)
	}
	return instance.ac.NewRequest(temp, op)
}

func (instance *Department) SetCapacity(cap int) {
	instance.employeesNum = cap
}

func GetDepartment(id int) *Department {
	instance := new(Department)
	instance.depId = id
	instance.employeesCame = 0
	instance.ac = ac.GetAc()
	return instance
}
