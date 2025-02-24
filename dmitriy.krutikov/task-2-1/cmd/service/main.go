package main

import (
	"fmt"
    "office-temperature/internal/temperature"
)

func main() {
    err := temperature.Run()
	if err != nil {
        fmt.Printf("Error: %s\n", err)
		return 
    }
}