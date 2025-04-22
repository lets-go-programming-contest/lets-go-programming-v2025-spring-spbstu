package process

import (
	"fmt"
	"errors"
	"github.com/conditioners/internal/streamUtils"
)

func getRequest() (int32, bool, error) {
	var (
		boundReqTypeSymb string
		isLowerBoundReq bool
		boundReq int32
	)
	
	boundReqTypeSymb, boundReq, err := streamUtils.ScanBoundReq()
	if err != nil {
		return 0, false, err
	}
	
	switch boundReqTypeSymb {
	case `>=`:
		isLowerBoundReq = true
	case `<=`:
		isLowerBoundReq = false
	default:
		return 0, false, errors.New("please use only '>='" +
					    "' and '<=' signs")
	}

	return boundReq, isLowerBoundReq, err
}

func getNewBounds(lowerBound int32, upperBound int32, boundReq int32,
	          isLowerBoundReq bool) (int32, int32) {
	if lowerBound == -1 {
		return lowerBound, upperBound
	}

        if isLowerBoundReq {
                if boundReq > lowerBound {
                        if boundReq <= upperBound {
                                lowerBound = boundReq
                        } else {
                                lowerBound = -1
                        }
                }
        } else {
                if boundReq < upperBound {
                        if boundReq >= lowerBound {
                                upperBound = boundReq
                        } else {
                                lowerBound = -1
                        }
                }
        }

        return lowerBound, upperBound
}

func processDepartment() error {
        var (
                lowerBound int32 = 15
                upperBound int32 = 30
        )

        nEmps, err := streamUtils.ScanUInt32()
	if err != nil {
		return err
	}

        for i := uint32(0) ; i < nEmps ; i += 1 {
		boundRec, isLowerBoundReq, err := getRequest()
		if err != nil {
			return err
		}

                lowerBound, upperBound = getNewBounds(lowerBound,
						      upperBound,
						      boundRec,
						      isLowerBoundReq)
		if err != nil {
			return err
		}

                fmt.Println(lowerBound)
        }

        return nil
}

func ProcessOffice() error {
        nDeps, err := streamUtils.ScanUInt32()
        if err != nil {
                return err
        }

        for i := uint32(0) ; i < nDeps ; i += 1 {
                err = processDepartment()
		if err != nil {
			return err
		}
        }

        return nil
}
