package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"

	"task-8/internal/debug"
	"task-8/internal/fib"
)

// Version set via ldflags; default is "dev"
var version = "dev"

//go:embed embedded/sample.txt
var sampleText embed.FS

func main() {
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Printf("Fibonacci(%d) = %d\n", n, fib.Fibonacci(n))
		} else {
			fmt.Println("Invalid argument, expected an integer.")
		}
	} else {
		fmt.Println("Fibonacci(10) =", fib.Fibonacci(10))
	}

	fmt.Println("Version:", version)

	// Read embedded file
	data, err := sampleText.ReadFile("embedded/sample.txt")
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
	} else {
		fmt.Println("Embedded file content:")
		fmt.Println(string(data))
	}

	// Debug output (active only with "debug" build tag)
	debug.Print()

	fmt.Println("Generated Fibonacci(10):", GeneratedFibonacci)
}
