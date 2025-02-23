package main

import (
	"flag"
	"log"

	"task-3/internal/config"
	"task-3/internal/fetcher"
	"task-3/internal/loader"
	"task-3/internal/output"
	"task-3/internal/processor"
)

func main() {
	// Define a flag for the configuration file path
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Load data (file or URL)
	data, err := loader.LoadData(cfg.InputFile)
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Fetch and parse data
	currencies, err := fetcher.FetchData(data)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	// Process data (sort by descending value)
	sortedCurrencies := processor.SortCurrencies(currencies)

	// Save result to JSON
	err = output.WriteToFile(cfg.OutputFile, sortedCurrencies)
	if err != nil {
		log.Fatalf("Failed to save result: %v", err)
	}

	log.Println("Processing completed successfully!")
}