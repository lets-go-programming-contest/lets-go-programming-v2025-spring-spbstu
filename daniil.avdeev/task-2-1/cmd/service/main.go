package main

import (
	"fmt"

	"github.com/realFrogboy/task-2-1/internal/conditioner"
	"github.com/realFrogboy/task-2-1/internal/parser"
)

func main() {
	n_commands, error := parser.GetInt()
	if error != nil {
		fmt.Println(error)
		return
	}
	if n_commands < 0 {
		fmt.Printf("invalid number of commands: %d", n_commands)
		return
	}

	for command := 0; command < n_commands; command++ {
		n_workers, error := parser.GetInt()
		if error != nil {
			fmt.Println(error)
			return
		}
		if n_workers < 0 {
			fmt.Printf("invalid number of workers: %d", n_workers)
			return
		}

		cond := conditioner.NewConditioner()
		for worker := 0; worker < n_workers; worker++ {
			TR, error := parser.GetTemperatureRegime()
			if error != nil {
				fmt.Println(error)
				return
			}

			T := cond.GetOptimalTemperature(&TR)
			fmt.Println(T)
		}
	}
}
