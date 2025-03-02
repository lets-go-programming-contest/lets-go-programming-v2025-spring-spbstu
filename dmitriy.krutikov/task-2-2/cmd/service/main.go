package main

import (
	"fmt"
	
	"choise/internal/choiceDishes"
)

func main() {
	err := choiseDishes.Run()
	
	if err != nil {
        fmt.Printf("Error: %s\n", err)
		return 
    }
}