package temperature_range

import "task-2-1/internal/input"

const (
	minTemperature = 15
	maxTemperature = 30
)

type TemperatureRange struct {
	min int
	max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		min: minTemperature,
		max: maxTemperature,
	}
}

func (tr *TemperatureRange) Update(op input.Operation, temp int) bool {
	switch op {
	case input.GreaterOrEqual:
		tr.min = max(tr.min, temp)
	case input.LessOrEqual:
		tr.max = min(tr.max, temp)
	}
	return tr.min <= tr.max
}

func (tr *TemperatureRange) GetOptimal() *int {
	if tr.min > tr.max {
		return nil
	}
	return &tr.min
}
