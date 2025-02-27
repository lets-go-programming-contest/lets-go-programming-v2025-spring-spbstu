package conditioner

const (
	MinTemperature = 15.0
	MaxTemperature = 30.0
)

type TemperatureRegime struct {
	Temp float32
	Reg  string
}

func average(arr []float32) float32 {
	if len(arr) == 0 {
		return 0.0
	}

	total := float32(0)
	for _, v := range arr {
		total += v
	}

	return total / float32(len(arr))
}

type Conditioner struct {
	lhsBorder float32
	rhsBorder float32
	temps     []float32
}

func NewConditioner() *Conditioner {
	return &Conditioner{MinTemperature, MaxTemperature, []float32{}}
}

func (cond *Conditioner) GetOptimalTemperature(TR *TemperatureRegime) float32 {
	switch TR.Reg {
	case "<=":
		if TR.Temp < cond.lhsBorder {
			return -1.0
		}

		if cond.rhsBorder > TR.Temp {
			cond.rhsBorder = TR.Temp
		}

	case ">=":
		if TR.Temp > cond.rhsBorder {
			return -1.0
		}

		if cond.lhsBorder < TR.Temp {
			cond.lhsBorder = TR.Temp
		}
	}

	cond.temps = append(cond.temps, TR.Temp)
	middle := average(cond.temps)

	if middle > cond.rhsBorder {
		return cond.rhsBorder
	}
	if middle < cond.lhsBorder {
		return cond.lhsBorder
	}
	return middle
}
