package interval

import "errors"

type IntervalValue struct {
	LowBound, UpBound int
	Value             int
}

func NewIntervalValue(low, up int) *IntervalValue {
	return &IntervalValue{
		LowBound: low,
		UpBound:  up,
	}
}

func (iv *IntervalValue) BiggerThan(low int) error {
	if iv.LowBound >= low {
		return nil
	}
	iv.LowBound = low

	if iv.Value >= low {
		return nil
	}
	iv.Value = low

	if iv.Value > iv.UpBound {
		iv.Value = -1
		return errors.New("Стали больше верхней границы")
	}
	return nil
}

func (iv *IntervalValue) LessThan(up int) error {
	if iv.UpBound <= up {
		return nil
	}
	iv.UpBound = up

	if iv.Value <= up {
		return nil
	}
	iv.Value = up

	if iv.Value < iv.LowBound {
		iv.Value = -1
		return errors.New("Стали меньше нижней границы")
	}
	return nil
}
