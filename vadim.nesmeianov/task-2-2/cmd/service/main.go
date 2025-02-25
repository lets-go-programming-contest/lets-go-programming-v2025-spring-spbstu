package main

import (
	"log"
	"task-2-2/internal/choice"
)

func main() {
	err := choice.Run()
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
