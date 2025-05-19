package office

import (
	"errors"
	"fmt"
)

func GetDepartment() (int, error) {
	var k int

	_, errScan := fmt.Scan(&k)
	if errScan != nil {
		return -1, errScan
	}

	if k < 1 || k > 1000 {
		err := errors.New("wrong employees number: it should be between 1 and 1000")
		return -1, err
	}

	return k, nil
}

func GetOffice() (int, error) {
	var n int

	_, errScan := fmt.Scan(&n)
	if errScan != nil {
		return -1, errScan
	}

	if n < 1 || n > 1000 {
		return -1, errors.New("wrong departments number: it should be between 1 and 1000")
	}

	return n, nil

}
