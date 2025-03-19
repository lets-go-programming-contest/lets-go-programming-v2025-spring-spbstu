package matrix

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LinearSystem struct {
	Coefs [][]float64 
	B      []float64  
}

func ReadSystemFromFile(path string) (*LinearSystem, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var system LinearSystem
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			return nil, fmt.Errorf("invalid equation format")
		}

		coefs := make([]float64, len(fields)-1)
		for i, f := range fields[:len(fields)-1] {
			val, err := strconv.ParseFloat(f, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", f)
			}
			coefs[i] = val
		}

		bVal, err := strconv.ParseFloat(fields[len(fields)-1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid RHS: %s", fields[len(fields)-1])
		}

		system.Coefs = append(system.Coefs, coefs)
		system.B = append(system.B, bVal)
	}

	return &system, nil
}

func (ls *LinearSystem) IsSquare() bool {
	return len(ls.Coefs) == len(ls.Coefs[0])
}

func (ls *LinearSystem) PrintLinearSystem() {
	for i := 0; i < len(ls.Coefs); i++ {
		for j := 0; j < len(ls.Coefs[0]); j++ {
			fmt.Printf("%v ", ls.Coefs[i][j])
		}
		fmt.Printf("| %v\n", ls.B[i])
	}
}