package main

import (
	"flag"
	"fmt"

	"github.com/yanelox/task-4/internal/mysum"
)

var GoroutinesCount = flag.Int("goroutines", 1, "Number of goroutines to run")
var Number = flag.Int("number", 1, "Number N")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Description: count sum of 1/k by k from 1 to N\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *GoroutinesCount > *Number {
		panic ("Goroutines count can't be more than number N")
	}

	fmt.Printf("Sum of 1/k by k from 1 to N is %f\n", mysum.CountSum(*Number, *GoroutinesCount))
}