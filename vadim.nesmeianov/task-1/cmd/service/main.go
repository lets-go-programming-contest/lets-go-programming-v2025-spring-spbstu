package main

import (
	"log"
	calc "task-1/internal/calc"
)

func main() {
	err := calc.Run()
	if err != nil {
		errorHandler(err)
	}
}

func errorHandler(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
