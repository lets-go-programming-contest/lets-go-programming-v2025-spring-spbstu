package main

import (
	"fmt"
	"sync"

	"github.com/kseniadobrovolskaia/task-4/internal/H2O"
)

func main() {
	const NumMolecules = 5
	var wg sync.WaitGroup
	wg.Add(NumMolecules * 3)

	Water := H2O.NewWater()

	fmt.Printf("Number of molecules: %d\n\t", NumMolecules)
	for i := 0; i < NumMolecules*3; i += 3 {
		go Water.RunHydrogen(i, &wg)
		go Water.RunHydrogen(i+1, &wg)
		go Water.RunOxygen(i+2, &wg)
	}

	wg.Wait()
}
