package optimal

import "task-2-1/internal/temperature"

func inInt(interval OptInt, t int) bool {
	return !(t < interval.T1 || t > interval.T2)
}

func GetOptInt(interval OptInt, t temperature.Temperature) OptInt {
	if inInt(interval, t.T) {
		if t.Less {
			return OptInt{interval.T1, t.T}
		} else {
			return OptInt{t.T, interval.T2}
		}
	} else {
		return OptInt{-1, -1}
	}
}
