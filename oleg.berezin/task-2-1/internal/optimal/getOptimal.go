package optimal

import (
	"task-2-1/internal/temperature"
)

func inInt(interval OptInt, t int) bool {
	return !(t < interval.T1 || t > interval.T2)
}

func isIntersected(inteval1 OptInt, t temperature.Temperature) bool {
	if t.Less {
		return inteval1.T2 < t.T
	} else {
		return inteval1.T1 > t.T
	}
}

func GetOptInt(interval OptInt, t temperature.Temperature) OptInt {
	if inInt(interval, t.T) {
		if t.Less {
			return OptInt{interval.T1, t.T}
		} else {
			return OptInt{t.T, interval.T2}
		}
	} else if isIntersected(interval, t) {
		return interval
	} else {
		return OptInt{-1, -1}
	}
}
