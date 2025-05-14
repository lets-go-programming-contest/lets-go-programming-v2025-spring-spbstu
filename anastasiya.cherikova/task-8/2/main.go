//go:build !debug

package main

import "fmt"

// The main version of the program
func main() {
	fmt.Println("Production Version:")
	result := calculate(10, 5)
	fmt.Printf("Result: %d\n", result)
}

func calculate(a, b int) int {
	return a * b
}
