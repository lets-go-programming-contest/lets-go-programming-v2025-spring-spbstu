//go:build debug

package main

import (
	"fmt"
	"time"
)

// Debug version of the program
func main() {
	fmt.Println("Debug Version:")
	start := time.Now()

	a, b := 10, 5
	fmt.Printf("Input values: a=%d, b=%d\n", a, b)

	result := calculate(a, b)
	fmt.Printf("Result: %d\n", result)

	fmt.Printf("Execution time: %v\n", time.Since(start))
}

func calculate(a, b int) int {
	fmt.Println("Performing calculation...")
	return a * b
}
