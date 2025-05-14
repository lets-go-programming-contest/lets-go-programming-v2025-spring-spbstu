package main

import (
	"bufio"
	"fmt"
	"os"
	"task-2-1/internal/input"
	"task-2-1/internal/temperature_range"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	processor := input.NewInputProcessor(scanner)

	departments, err := processor.ReadDepartments()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, dept := range departments {
		tr := temperature_range.NewTemperatureRange()
		for _, req := range dept.Requests {
			ok := tr.Update(req.Op, req.Temp)
			if !ok {
				fmt.Println(-1)
				continue
			}
			optimalTemp := tr.GetOptimal()
			if optimalTemp == nil {
				fmt.Println(-1)
			} else {
				fmt.Println(*optimalTemp)
			}
		}
	}
}
