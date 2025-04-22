package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/denisboborukhin/currency_converter/internal/config"
	"github.com/denisboborukhin/currency_converter/internal/currency"
	"github.com/denisboborukhin/currency_converter/internal/file"
)

func main() {
	configFile := flag.String("config", "", "Path to the config file")
	flag.Parse()

	if *configFile == "" {
		log.Fatal("Config file path is required")
	}

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	currenciesJSON, err := currency.ConvertToJSON(cfg.InputFile)
	if err != nil {
		log.Fatalf("Error converting currencies: %v", err)
	}

	currency.SortCurrenciesByValue(currenciesJSON)

	err = file.SaveCurrencies(cfg.OutputFile, currenciesJSON)
	if err != nil {
		log.Fatalf("Error saving currencies: %v", err)
	}

	fmt.Println("Currencies processed and saved successfully")
}
