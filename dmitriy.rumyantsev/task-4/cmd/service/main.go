package main

import (
	"github.com/dmitriy.rumyantsev/task-4/internal/simulator"
)

func main() {
	sim := simulator.NewBankSimulator()

	sim.RunBothSimulations()
}
