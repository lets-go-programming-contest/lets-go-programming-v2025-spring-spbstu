package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/realFrogboy/task-3/internal/configparser"
	"github.com/realFrogboy/task-3/internal/cursesparser"
	"github.com/realFrogboy/task-3/internal/jsonemitter"
)

func main() {
	configFilePath := flag.String("config", "", "Path to the config file")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		fmt.Printf("main(): can't read config file: %s", err)
		return
	}

	config, err := configparser.Parse(configFile)
	if err != nil {
		fmt.Printf("main(): can't parse config file: %s", err)
		return
	}

	cursesFile, err := os.Open(config.InputFile)
	if err != nil {
		fmt.Printf("main(): can't open input file: %s", err)
		return
	}
	defer cursesFile.Close()

	valutes, err := cursesparser.Parse(cursesFile)
	if err != nil {
		fmt.Printf("main(): cant parse input file: %s", err)
		return
	}

	outputFile, err := os.OpenFile(config.OutputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("main(): can't open output file: %s", err)
		return
	}
	defer outputFile.Close()

	err = jsonemitter.Emit(outputFile, valutes)
	if err != nil {
		fmt.Printf("main(): can't emit valutes: %s", err)
		return
	}
}
