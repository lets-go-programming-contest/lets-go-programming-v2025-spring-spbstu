package main

import (
	"flag"
	"fmt"
	"os"

	"task-3/internal/decoder"
	"task-3/internal/encoder"
	"task-3/internal/manager"
	"task-3/internal/transformer"
)

func process() error {
	configPath := flag.String("config", "config.yam", "Path to configuration file")
	flag.Parse()
	config, err := manager.ParseConfig(*configPath)
	if err != nil {
		return fmt.Errorf("bad parse: %v", err)
	}
	rawData, err := decoder.Decode(config.Input)
	if err != nil {
		return fmt.Errorf("bad decode: %v", err)
	}
	data, err := transformer.Transform(rawData)
	if err != nil {
		return fmt.Errorf("bad transform: %v", err)
	}
	err = encoder.Encode(data, config.Output)
	if err != nil {
		return fmt.Errorf("bad encode: %v", err)
	}
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
