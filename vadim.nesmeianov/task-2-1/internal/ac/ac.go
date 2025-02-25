package ac

import (
	"errors"
	"task-2-1/internal/comp_op"
)

type AcRequest struct {
	temp int
	op   comp_op.Operand
}

type Ac struct {
	maxTemp     int
	minTemp     int
	currentTemp int
	requests    []AcRequest
}

func (instance *Ac) processRequest(req AcRequest) error {
	instance.requests = append(instance.requests, req)

	if instance.currentTemp != -1 {
		minNewTemp := instance.minTemp
		maxNewTemp := instance.maxTemp

		for _, rq := range instance.requests {
			switch rq.op {
			case comp_op.BiggerOrEqual:
				res, err := comp_op.Compare(rq.op, rq.temp, minNewTemp)
				if err != nil {
					return err
				}
				if res {
					minNewTemp = rq.temp
				}

			case comp_op.LessOrEqual:
				res, err := comp_op.Compare(rq.op, rq.temp, maxNewTemp)
				if err != nil {
					return err
				}
				if res {
					maxNewTemp = rq.temp
				}

			default:
				return errors.New("unknown operator")
			}
		}

		if minNewTemp <= maxNewTemp {
			instance.currentTemp = minNewTemp
		} else {
			instance.currentTemp = -1
		}
	}

	return nil
}

func (instance *Ac) NewRequest(temp int, op comp_op.Operand) error {
	req := AcRequest{temp, op}
	return instance.processRequest(req)
}

func (instance Ac) GetTemperature() int {
	return instance.currentTemp
}

func GetAc() Ac {
	const minTemp = 15
	const maxTemp = 30
	var instance Ac

	instance.maxTemp = maxTemp
	instance.minTemp = minTemp
	instance.currentTemp = 0
	instance.requests = make([]AcRequest, 0, 2)
	instance.requests = append(instance.requests, AcRequest{minTemp, comp_op.BiggerOrEqual})
	instance.requests = append(instance.requests, AcRequest{maxTemp, comp_op.LessOrEqual})

	return instance
}
