package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"

	"github.com/kseniadobrovolskaia/task-3/internal/config"
	"github.com/kseniadobrovolskaia/task-3/internal/sortValue"
	"github.com/kseniadobrovolskaia/task-3/internal/xmlToJson"
)

var configPath = flag.String("config", "", "Path to config file .yaml")

func main() {
	flag.Parse()
	if *configPath == "" {
		color.Red("--config not specified")
		os.Exit(0)
	}

	// Read config file
	config, err := config.ReadConfigFile(*configPath)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	// Read XML file
	vals, err := xmlToJson.ReadXMLFile(config.InputFile)
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}

	// Sorting
	sort.Sort(sortValue.ByValue(vals))

	// Dump in Json file
	lenBytes, err := xmlToJson.WriteInJSONFile(config.OutputFile, vals)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d %s\n", lenBytes, color.GreenString("bytes written in "+config.OutputFile))
}
