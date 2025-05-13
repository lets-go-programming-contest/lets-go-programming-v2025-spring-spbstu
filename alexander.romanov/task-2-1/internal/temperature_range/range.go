package temperature_range

import "task-2-1/internal/input"

const (
	MinTemperature = 15
	MaxTemperature = 30
)

type TemperatureRange struct {
	Min int
	Max int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		Min: MinTemperature,
		Max: MaxTemperature,
	}
}

func (tr *TemperatureRange) Update(op input.Operation, temp int) bool {
	switch op {
	case input.GreaterOrEqual:
		tr.Min = max(tr.Min, temp)
	case input.LessOrEqual:
		tr.Max = min(tr.Max, temp)
	}
	return tr.Min <= tr.Max
}

func (tr *TemperatureRange) GetOptimal() *int {
	if tr.Min > tr.Max {
		return nil
	}
	return &tr.Min
}
