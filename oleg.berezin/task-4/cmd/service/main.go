package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"task-4/internal/wave"
)

func main() {
	length := 1.0
	time := 2.0
	threadNum := 8

	nx := int(length / wave.Dx)
	nt := int(time / wave.Dt)

	u := wave.WaveSim(nx, nt, threadNum)

	file, err := os.Create("wave_output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for t := 0; t < nt; t++ {
		row := make([]string, nx)
		for i := 0; i < nx; i++ {
			row[i] = strconv.FormatFloat(u[t][i], 'f', 6, 64)
		}
		writer.Write(row)
	}
}
