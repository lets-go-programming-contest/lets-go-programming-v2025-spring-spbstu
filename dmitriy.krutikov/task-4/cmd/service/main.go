package main

import (
	"flag"
	"fmt"
	"linear-system/internal/jacobi"
	"linear-system/internal/matrix"
	"os"
)

func main() {
	filePath := flag.String("f", "equations.txt", "Path to equations file")
	useSync := flag.Bool("sync", true, "Use synchronized version")
	maxIter := flag.Int("iter", 1000, "Maximum iterations")
	tolerance := flag.Float64("tol", 1e-6, "Tolerance for convergence")
	flag.Parse()

	system, err := matrix.ReadSystemFromFile(*filePath)
	if err != nil {
		fmt.Printf("Error reading system: %v\n", err)
		os.Exit(1)
	}

	if !system.IsSquare() {
		fmt.Println("Matrix must be square")
		os.Exit(1)
	}

	//system.PrintLinearSystem()
	var solution []float64
	
	if *useSync {
		solution = jacobi.SolveSync(system, *maxIter, *tolerance)
	} else {
		solution = jacobi.SolveAsync(system, *maxIter, *tolerance)
	}

	fmt.Println("Solution:")
	for i, val := range solution {
		fmt.Printf("x%d = %.6f\n", i+1, val)
	}
}