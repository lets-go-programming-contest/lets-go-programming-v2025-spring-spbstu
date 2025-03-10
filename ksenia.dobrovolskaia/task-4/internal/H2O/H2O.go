package H2O

import (
	"fmt"
	"sync"
)

type Water struct {
	hydrogens int
	oxydens   int
	mtx       sync.Mutex
	cond      sync.Cond
}

func NewWater() *Water {
	w := &Water{
		mtx: sync.Mutex{},
	}

	w.cond = *sync.NewCond(&w.mtx)
	return w
}

func (w *Water) RunHydrogen(Id int, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mtx.Lock()
	defer w.mtx.Unlock()

	for w.hydrogens >= 2 {
		w.cond.Wait()
	}
	fmt.Printf("H(%d)", Id)
	w.hydrogens++

	if w.hydrogens >= 2 && w.oxydens >= 1 {
		w.releaseMolecule()
	}

	w.cond.Broadcast()
}

func (w *Water) RunOxygen(Id int, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mtx.Lock()
	defer w.mtx.Unlock()

	for w.oxydens >= 1 {
		w.cond.Wait()
	}
	fmt.Printf("O(%d)", Id)
	w.oxydens++

	if w.hydrogens >= 2 {
		w.releaseMolecule()
	}

	w.cond.Broadcast()
}

func (w *Water) releaseMolecule() {
	w.oxydens -= 1
	w.hydrogens -= 2
	fmt.Printf("\n\t")
}
