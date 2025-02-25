package ac

import (
	"errors"
	"task-2-1/pkg/comp_op"
)

type AcRequest struct {
	temp int
	op   comp_op.Operand
}

type ac struct {
	maxTemp     int
	minTemp     int
	currentTemp int
	requests    []AcRequest
}

type Ac interface {
	NewRequest(temp int, op comp_op.Operand) error
	GetTemperature() int
}

func (instance *ac) processRequest(req AcRequest) error {
	instance.requests = append(instance.requests, req)

	if instance.currentTemp != -1 {
		minNewTemp := instance.minTemp
		maxNewTemp := instance.maxTemp

		for _, rq := range instance.requests {
			switch rq.op {
			case comp_op.BiggerOrEqual:
				if rq.temp >= minNewTemp {
					minNewTemp = rq.temp
				}

			case comp_op.LessOrEqual:
				if rq.temp <= maxNewTemp {
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

func (instance *ac) NewRequest(temp int, op comp_op.Operand) error {
	req := AcRequest{temp, op}
	return instance.processRequest(req)
}

func (instance ac) GetTemperature() int {
	return instance.currentTemp
}

func GetAc() Ac {
	const minTemp = 15
	const maxTemp = 30

	instance := new(ac)
	instance.maxTemp = maxTemp
	instance.minTemp = minTemp
	instance.currentTemp = 0
	instance.requests = make([]AcRequest, 0, 2)
	instance.requests = append(instance.requests, AcRequest{minTemp, comp_op.BiggerOrEqual})
	instance.requests = append(instance.requests, AcRequest{maxTemp, comp_op.LessOrEqual})

	return instance
}
