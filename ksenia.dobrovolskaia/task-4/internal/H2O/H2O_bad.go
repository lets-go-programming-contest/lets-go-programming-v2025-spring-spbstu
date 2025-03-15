package H2O

import (
	"fmt"
	"sync"
)

type WaterBad struct {
	hydrogens int
	oxydens   int
}

func NewWaterBad() *WaterBad {
	return &WaterBad{}
}

func (w *WaterBad) RunHydrogen(Id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("H(%d)", Id)
	w.hydrogens++

	if w.hydrogens >= 2 && w.oxydens >= 1 {
		w.releaseMolecule()
	}
}

func (w *WaterBad) RunOxygen(Id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("O(%d)", Id)
	w.oxydens++

	if w.hydrogens >= 2 {
		w.releaseMolecule()
	}
}

func (w *WaterBad) releaseMolecule() {
	w.oxydens -= 1
	w.hydrogens -= 2
	fmt.Printf("\n\t")
}
