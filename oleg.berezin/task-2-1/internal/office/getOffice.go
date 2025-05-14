package office

import (
	"errors"
	"fmt"
)

func getDepartment() (Department, error) {
	var k int

	_, errScan := fmt.Scan(&k)
	if errScan != nil {
		return Department{}, errScan
	}

	if k < 1 || k > 1000 {
		err := errors.New("Wrong employees number: it should be between 1 and 1000")
		return Department{}, err
	}

	return Department{k}, nil
}

func getOffice() (Office, error) {
	var n int

	_, errScan := fmt.Scan(&n)
	if errScan != nil {
		return Office{}, errScan
	}

	if n < 1 || n > 1000 {
		return Office{}, errors.New("Wrong departments number: it should be between 1 and 1000")
	}

	dept, err := getDepartment()
	if err != nil {
		return Office{}, err
	}

	return Office{
		DepartmentsCount: n,
		EmployeesCount:   dept.EmployeesCount,
	}, nil
}
