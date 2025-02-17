package process

import (
	"fmt"
	"errors"
)

func getRequest() (int32, bool, error) {
	var (
		boundReqTypeSymb string
		isLowerBoundReq bool
		boundReq int32
	)
	
	_, err := fmt.Scan(&boundReqTypeSymb, &boundReq)
	if err != nil {
		return 0, false, err
	}
	
	switch boundReqTypeSymb {
	case `>=`:
		isLowerBoundReq = true
	case `<=`:
		isLowerBoundReq = false
	default:
		err = errors.New("please use only '>=' and '<=' signs")
	}

	return boundReq, isLowerBoundReq, err
}

func getNewBounds(lowerBound int32, upperBound int32, boundReq int32,
	          isLowerBoundReq bool) (int32, int32, error) {
	if lowerBound == -1 {
		return lowerBound, upperBound, nil
	}

        if isLowerBoundReq {
                if boundReq > lowerBound {
                        if boundReq < upperBound {
                                lowerBound = boundReq
                        } else {
                                lowerBound = -1
                        }
                }
        } else {
                if boundReq < upperBound {
                        if boundReq > lowerBound {
                                upperBound = boundReq
                        } else {
                                lowerBound = -1
                        }
                }
        }

        return lowerBound, upperBound, nil
}

func processDepartment() error {
        var (
                nEmps uint32
                lowerBound int32 = 15
                upperBound int32 = 30
        )

        _, err := fmt.Scan(&nEmps)
	if err != nil {
		return err
	}

        var i uint32
        for i = 0 ; i < nEmps ; i += 1 {
		boundRec, isLowerBoundReq, err := getRequest()
		if err != nil {
			break
		}

                lowerBound, upperBound, err = getNewBounds(lowerBound,
							   upperBound,
							   boundRec,
							   isLowerBoundReq)
		if err != nil {
			break
		}

                fmt.Println(lowerBound)
        }

        return err
}

func ProcessOffice() error {
        var nDeps uint32

        _, err := fmt.Scan(&nDeps)
        if err != nil {
                return err
        }

        var i uint32
        for i = 0 ; i < nDeps ; i += 1 {
                err = processDepartment()
        }

        return err
}
