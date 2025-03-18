package jacobi

import (
	"linear-system/internal/matrix"
	"math"
	"math/rand"
	"sync"
	"time"
)

func SolveSync(system *matrix.LinearSystem, maxIter int, tol float64) []float64 {
	n := len(system.B)
	x := make([]float64, n)
	xNew := make([]float64, n)
	var wg sync.WaitGroup

	for iter := 0; iter < maxIter; iter++ {
		wg.Add(n)
		
		for i := 0; i < n; i++ {
			go updateSync(i, system, x, xNew, &wg)
		}
		
		wg.Wait()
		x, xNew = xNew, x 
		
		if checkConvergence(x, xNew, tol) {
			break
		}
	}
	return x
}

func SolveAsync(system *matrix.LinearSystem, maxIter int, tol float64) []float64 {
	n := len(system.B)
	x := make([]float64, n)
	xNew := make([]float64, n)

	for iter := 0; iter < maxIter; iter++ {
		for i := 0; i < n; i++ {
			go updateAsync(i, system, x, xNew)
		}
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		x, xNew = xNew, x 
		
		if checkConvergence(x, xNew, tol) {
			break
		}
	}
	return x
}


func checkConvergence(old, new []float64, tol float64) bool {
	maxDiff := 0.0
	for i := range old {
		diff := math.Abs(old[i] - new[i])
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff < tol
}
