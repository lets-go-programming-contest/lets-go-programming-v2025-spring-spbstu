package main

import (
	"fmt"
	"task-2-1/internal/office"
	"task-2-1/internal/optimal"
	"task-2-1/internal/temperature"
)

func main() {
	Y, err := office.GetOffice()
	if err != nil {
		fmt.Printf("error during GetOffice\n")
		panic(err)
	}

	for n := 0; n < Y; n++ {
		n, err := office.GetDepartment()
		if err != nil {
			fmt.Printf("error during GetDepartment\n")
			panic(err)
		}

		t0, err := temperature.GetTemperature()
		if err != nil {
			fmt.Printf("error during GetTemperature\n")
			panic(err)
		}

		optInt := optimal.GetOptInt(optimal.OptInt{T1: 15, T2: 30}, t0)

		optimal.ShowOptimal(optInt)

		for i := 1; i < n; i++ {
			t, err := temperature.GetTemperature()
			if err != nil {
				fmt.Printf("error during GetTemperature\n")
				panic(err)
			}

			optInt = optimal.GetOptInt(optInt, t)

			optimal.ShowOptimal(optInt)
		}
	}
}
