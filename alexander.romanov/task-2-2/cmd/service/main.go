package main

import (
	"fmt"
	"os"
	"task-2-2/internal/input"
	"task-2-2/internal/preference"
)

func main() {
	data, err := input.ReadInput(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	result := preference.GetKthPrefered(data.Preferences, data.K)

	fmt.Println(result)
}
