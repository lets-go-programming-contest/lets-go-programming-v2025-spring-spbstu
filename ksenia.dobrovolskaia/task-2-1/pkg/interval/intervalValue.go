package interval

import "errors"

//	LowBound   Value  UpBound
//
// ------|---------|-----|-------
type IntervalValue struct {
	isInitLow, isInitUp bool
	LowBound, UpBound   int
	Value               int
}

// BiggerThan увеличивает iv.Value до low, если iv.Value < low.
func (iv *IntervalValue) BiggerThan(low int) error {
	if !iv.isInitLow && !iv.isInitUp {
		iv.LowBound = low
		iv.isInitLow = true
		iv.Value = low
		return nil
	} else if iv.isInitLow && iv.LowBound >= low {
		// Требование уже удовлетворено
		return nil
	}
	iv.isInitLow = true
	iv.LowBound = low

	// Требование уже удовлетворено
	if iv.Value >= low {
		return nil
	}
	iv.Value = low

	// Стали больше верхней границы -> несуществующее состояние
	if iv.isInitUp && iv.Value > iv.UpBound {
		iv.Value = -1
		return errors.New("Стали больше верхней границы")
	}
	return nil
}

// LessThan уменьшает iv.Value до up, если iv.Value > up.
func (iv *IntervalValue) LessThan(up int) error {
	if !iv.isInitLow && !iv.isInitUp {
		iv.UpBound = up
		iv.isInitUp = true
		iv.Value = up
		return nil
	} else if iv.isInitUp && iv.UpBound <= up {
		// Требование уже удовлетворено
		return nil
	}
	iv.isInitUp = true
	iv.UpBound = up

	// Требование уже удовлетворено
	if iv.isInitLow && iv.Value <= up {
		return nil
	}
	iv.Value = up

	// Стали меньше нижней границы -> несуществующее состояние
	if iv.Value < iv.LowBound {
		iv.Value = -1
		return errors.New("Стали меньше нижней границы")
	}
	return nil
}
