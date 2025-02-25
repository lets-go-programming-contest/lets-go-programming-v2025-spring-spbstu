package main

import (
	"log"
	"task-2-1/internal/office"
)

func main() {
	err := office.Run(true)

	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
