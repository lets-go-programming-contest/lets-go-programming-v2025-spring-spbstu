package jacobi

import (
	"linear-system/internal/matrix"
	"math/rand"
	"sync"
	"time"
)

func updateSync(i int, system *matrix.LinearSystem, x, xNew []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0.0
	diag := system.Coefs[i][i]
	
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	for j := 0; j < len(x); j++ {
		if j != i {
			sum += system.Coefs[i][j] * x[j]
		}
	}

	xNew[i] = (system.B[i] - sum) / diag
}

func updateAsync(i int, system *matrix.LinearSystem, x, xNew []float64) {
	sum := 0.0
	diag := system.Coefs[i][i]
	
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

	for j := 0; j < len(x); j++ {
		if j != i {
			sum += system.Coefs[i][j] * x[j]
		}
	}

	xNew[i] = (system.B[i] - sum) / diag
}