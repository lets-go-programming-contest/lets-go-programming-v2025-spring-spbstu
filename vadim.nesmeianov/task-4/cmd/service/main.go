package main

import (
	"fmt"
	cliparser "integrate/internal/cli_parser"
	"integrate/pkg/integrator"
	"math"
)

func main() {

	nThreads, a, b, accuracyPow := cliparser.ReadArguments()

	f := func(x float64) float64 {
		if x == 0 {
			return 0
		} else {
			return math.Sin(1.0 / x)
		}
	}

	result := integrator.Integrate(a, b, accuracyPow, f, nThreads)

	fmt.Println("S(sin(1/x))dx = ", result, " from ", a, " to ", b)
	fmt.Println("Threads: ", nThreads)
	fmt.Println("Accuracy Power: 10^(", accuracyPow, ")")
}
