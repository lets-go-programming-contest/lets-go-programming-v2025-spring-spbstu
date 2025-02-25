package parser

import (
	"fmt"
	"strings"

	"github.com/realFrogboy/task-2-1/internal/conditioner"
)

func GetInt() (int, error) {
	var val int

	_, error := fmt.Scanf("%d", &val)
	if error != nil {
		return 0, error
	}

	return val, nil
}

func GetTemperatureRegime() (conditioner.TemperatureRegime, error) {
	var regime string
	var temperature float32

	_, error := fmt.Scanf("%s %f", &regime, &temperature)
	if error != nil {
		return conditioner.TemperatureRegime{}, error
	}
	regime = strings.TrimSpace(regime)

	if temperature < conditioner.MinTemperature || temperature > conditioner.MaxTemperature {
		return conditioner.TemperatureRegime{}, fmt.Errorf("invalid temperature: %f (should take [%f, %f])", temperature, conditioner.MinTemperature, conditioner.MaxTemperature)
	}

	if regime != ">=" && regime != "<=" {
		return conditioner.TemperatureRegime{}, fmt.Errorf("invalid regime: %q (can be >= or <=)", regime)
	}

	return conditioner.TemperatureRegime{temperature, regime}, nil
}
