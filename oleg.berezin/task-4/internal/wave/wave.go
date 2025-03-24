package wave

import (
	"math"
	"sync"
)

func WaveSim(nx int, nt int, threadNum int) [][]float64 {
	r := (C * Dt / Dx) * (C * Dt / Dx)

	u := make([][]float64, nt)
	for i := range u {
		u[i] = make([]float64, nx)
	}

	// initial
	for i := range nx {
		x := float64(i) * Dx
		u[0][i] = math.Exp(-100 * (x - 0.5) * (x - 0.5))
	}

	// second layer
	for i := 1; i < nx-1; i++ {
		u[1][i] = u[0][i] + 0.5*r*(u[0][i+1]-2*u[0][i]+u[0][i-1])
	}

	// main compute
	for t := 1; t < nt-1; t++ {
		var wg sync.WaitGroup
		wg.Add(threadNum)

		chunkSize := (nx - 2) / threadNum

		for w := range threadNum {
			start := 1 + w*chunkSize
			end := start + chunkSize
			if w == threadNum-1 {
				end = nx - 1
			}

			go func(start, end int) {
				defer wg.Done()

				for i := start; i < end; i++ {
					u[t+1][i] = 2*u[t][i] - u[t-1][i] + r*(u[t][i+1]-2*u[t][i]+u[t][i-1])
				}
			}(start, end)
		}
		wg.Wait()
	}

	return u
}
