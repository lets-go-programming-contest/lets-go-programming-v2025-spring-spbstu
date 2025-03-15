package main

import (
	"flag"

	"currency-converter/internal/config"
	"currency-converter/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("Config path is required")
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}


	curr, err := currency.Decode(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortCurrencies(curr)

	err = currency.SaveToJSON(curr, cfg.OutputFile)
	if err != nil {
		panic(err)
	}

}