package wave

import (
	"math"
	"sync"
)

type wave struct {
	u  [][]float64
	wg sync.WaitGroup
}

func constWave(nt, nx int) *wave {
	w := &wave{
		u: make([][]float64, nt),
	}
	for i := range w.u {
		w.u[i] = make([]float64, nx)
	}
	return w
}

func WaveSim(nx int, nt int, threadNum int) [][]float64 {
	r := (C * Dt / Dx) * (C * Dt / Dx)

	wave := constWave(nt, nx)

	// initial
	for i := range nx {
		x := float64(i) * Dx
		wave.u[0][i] = math.Exp(-100 * (x - 0.5) * (x - 0.5))
	}

	// second layer
	for i := 1; i < nx-1; i++ {
		wave.u[1][i] = wave.u[0][i] + 0.5*r*(wave.u[0][i+1]-2*wave.u[0][i]+wave.u[0][i-1])
	}

	// main compute
	for t := 1; t < nt-1; t++ {
		wave.wg.Add(threadNum)

		chunkSize := (nx - 2) / threadNum

		for w := range threadNum {
			start := 1 + w*chunkSize
			end := start + chunkSize
			if w == threadNum-1 {
				end = nx - 1
			}

			go func(start, end int) {
				defer wave.wg.Done()

				for i := start; i < end; i++ {
					wave.u[t+1][i] = 2*wave.u[t][i] - wave.u[t-1][i] + r*(wave.u[t][i+1]-2*wave.u[t][i]+wave.u[t][i-1])
				}
			}(start, end)
		}
		wave.wg.Wait()
	}

	return wave.u
}
