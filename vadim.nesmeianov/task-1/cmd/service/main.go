package main

import (
	"log"
	calc "task-1/internal/calc"
)

func main() {
	err := calc.Run()
	if err != nil {
		log.Fatal(err)
	}
}
