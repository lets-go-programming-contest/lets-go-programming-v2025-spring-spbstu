package main

import (
	"fmt"

	"github.com/realFrogboy/task-2-1/internal/conditioner"
	"github.com/realFrogboy/task-2-1/internal/parser"
)

func main() {
	nCommands, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if nCommands < 0 {
		fmt.Printf("invalid number of commands: %d", nCommands)
		return
	}

	for command := 0; command < nCommands; command++ {
		nWorkers, error := parser.GetInt()
		if error != nil {
			fmt.Println(error)
			return
		}
		if nWorkers < 0 {
			fmt.Printf("invalid number of workers: %d", nWorkers)
			return
		}

		cond := conditioner.NewConditioner()
		for worker := 0; worker < nWorkers; worker++ {
			temperatureRegime, error := parser.GetTemperatureRegime()
			if error != nil {
				fmt.Println(error)
				return
			}

			optimalTemperature := cond.GetOptimalTemperature(&temperatureRegime)
			fmt.Println(optimalTemperature)
		}
	}
}
