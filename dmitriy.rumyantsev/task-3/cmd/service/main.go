package main

import (
	"flag"
	"fmt"

	"github.com/dmitriy.rumyantsev/task-3/internal/config"
	"github.com/dmitriy.rumyantsev/task-3/internal/parser"
	"github.com/dmitriy.rumyantsev/task-3/internal/writer"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Parse command line arguments
	configPath := flag.String("config", "./configs/default_config.yaml", "Path to YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("You must specify the path to the configuration file (use -config flag)")
	}

	validate := validator.New()

	// Read and parse configuration file
	cfg, err := config.ReadConfig(*configPath, validate)
	if err != nil {
		panic(fmt.Sprintf("Configuration error: %v", err))
	}

	// Read and parse XML file with currency data
	valCurs, err := parser.ReadValCurs(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error reading currency data: %v", err))
	}

	// Process and sort currency data
	outputValutes, err := parser.ProcessValutesData(valCurs)
	if err != nil {
		panic(fmt.Sprintf("Error processing currency data: %v", err))
	}

	// Save processed data as JSON
	err = writer.SaveAsJSON(outputValutes, cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Error saving to output file: %v", err))
	}

	fmt.Println("Currency data successfully processed and saved to", cfg.OutputFile)
}
