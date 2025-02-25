package parser

import (
	"fmt"
	"strings"

	"github.com/realFrogboy/task-2-1/internal/conditioner"
)

func GetInt() (int, error) {
	var Val int

	_, error := fmt.Scanf("%d", &Val)
	if error != nil {
		return 0, error
	}

	return Val, nil
}

func GetTemperatureRegime() (conditioner.TemperatureRegime, error) {
	var R string
	var T float32

	_, error := fmt.Scanf("%s %f", &R, &T)
	if error != nil {
		return conditioner.TemperatureRegime{}, error
	}
	R = strings.TrimSpace(R)

	if T < conditioner.MinTemperature || T > conditioner.MaxTemperature {
		return conditioner.TemperatureRegime{}, fmt.Errorf("invalid temperature: %f (should take [%f, %f])", T, conditioner.MinTemperature, conditioner.MaxTemperature)
	}

	if R != ">=" && R != "<=" {
		return conditioner.TemperatureRegime{}, fmt.Errorf("invalid regime: %q (can be %q or %q)", R, ">=", "<=")
	}

	return conditioner.TemperatureRegime{T, R}, nil
}
