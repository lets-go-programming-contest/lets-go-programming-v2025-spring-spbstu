package main

import (
	"flag"
	"log"

	"github.com/vktr-ktzv/task3/internal/configLoader"
	"github.com/vktr-ktzv/task3/internal/dataHandler"
	"github.com/vktr-ktzv/task3/internal/output"
	"github.com/vktr-ktzv/task3/internal/parser"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()

	config := configLoader.LoadConfig(configPath)

	var err error
	valCurs, err := parser.ParseXML(config.InputFile)

	if err != nil {
		log.Fatal(err)
	}

	currencies := dataHandler.ConvertAndSort(valCurs)

	output.SaveJSON(config.OutputFile, currencies)
}
