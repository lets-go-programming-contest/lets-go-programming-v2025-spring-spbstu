package temperature

type TemperatureRange struct {
	effectiveMin int // max from all >=
	effectiveMax int // min from all <=
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		effectiveMin: 15, // default min temp
		effectiveMax: 30, // default max temp
	}
}

// ApplyConstraint applies new constrait and returns current optimal value
func (tr *TemperatureRange) ApplyConstraint(constraintType string, value int) int {
	switch constraintType {
	case ">=":
		if value > tr.effectiveMin {
			tr.effectiveMin = value
		}
	case "<=":
		if value < tr.effectiveMax {
			tr.effectiveMax = value
		}
	}

	currentMin := max(tr.effectiveMin, 15)
	currentMax := min(tr.effectiveMax, 30)

	// check constraits
	if currentMin > currentMax {
		return -1
	}
	
	// Optimal temperature - minimal from valid range (as pointed in task example)
	return currentMin
}

func max(a, b int) int { if a > b { return a }; return b }
func min(a, b int) int { if a < b { return a }; return b }