package cliparser

import "flag"

func ReadArguments() (int, float64, float64, int) {
	nThreads := flag.Int("n", 1, "Threads number")
	begin := flag.Float64("begin", 0, "Begin point of integration interval")
	end := flag.Float64("end", 1, "Endpoint of integration interval")
	accuracy := flag.Int("e", -16, "Accuracy: 10^e")

	flag.Parse()

	return *nThreads, *begin, *end, *accuracy
}
