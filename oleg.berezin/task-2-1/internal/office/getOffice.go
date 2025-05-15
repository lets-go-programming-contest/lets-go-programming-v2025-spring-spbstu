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
		err := errors.New("wrong employees number: it should be between 1 and 1000")
		return Department{}, err
	}

	return Department{k}, nil
}

func GetOffice() (Office, error) {
	var n int

	_, errScan := fmt.Scan(&n)
	if errScan != nil {
		return Office{}, errScan
	}

	if n < 1 || n > 1000 {
		return Office{}, errors.New("wrong departments number: it should be between 1 and 1000")
	}

	departments := make([]Department, 0, n)
	for range n {
		dept, err := getDepartment()
		if err != nil {
			return Office{}, err
		}
		departments = append(departments, dept)
	}

	return Office{Departments: departments}, nil

}
