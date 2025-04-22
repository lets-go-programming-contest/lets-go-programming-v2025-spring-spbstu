package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"task-3/internal/config"
	"task-3/internal/currency"
	"task-3/internal/jsonwriter"
	"task-3/internal/xmlparser"
)

func main() {

	// Add the -config flag with the default value
	configPath := flag.String(
		"config",
		"internal/config/config.yaml", // Default path
		"Path to configuration file",
	)
	flag.Parse()

	// Recover from any unexpected panics in the application
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Critical error: %v", r)
			os.Exit(1)
		}
	}()

	// Load configuration from config/config.yaml
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		handleError(err)
	}

	// Parse XML file with currency data
	currencies, err := xmlparser.ParseXML(cfg.InputFile)
	if err != nil {
		handleError(err)
	}

	// Sort currencies in descending order by value
	currency.SortCurrencies(currencies)

	// Write sorted results to JSON file
	if err := jsonwriter.WriteJSON(currencies, cfg.OutputFile); err != nil {
		handleError(err)
	}

	fmt.Printf("Success! Output saved to: %s\n", cfg.OutputFile)
}

// handleError processes application errors and exits the program.
// It unwraps joined errors and prints individual error messages.
func handleError(err error) {
	var joinedErr interface{ Unwrap() []error }

	// Check if error contains multiple wrapped errors
	if errors.As(err, &joinedErr) {
		for _, e := range joinedErr.Unwrap() {
			log.Printf("Error: %v", e)
		}
	} else {
		log.Printf("Error: %v", err)
	}
	os.Exit(1)
}
