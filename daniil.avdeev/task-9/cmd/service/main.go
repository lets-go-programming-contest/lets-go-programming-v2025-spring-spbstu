package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/realFrogboy/task-9/internal/config"
	"github.com/realFrogboy/task-9/internal/handler"
)

func readConfig(path string) config.Config {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read config file: %s", err)
	}

	config, err := config.Parse(configFile)
	if err != nil {
		log.Fatalf("Can't parse config file: %s", err)
	}

	return config
}

func main() {
	configFilePath := flag.String("config", "configs/config.yaml", "Path to the config file")
	flag.Parse()

	config := readConfig(*configFilePath)

	contactHandler, err := handler.NewContactHandler(config.DBPath)
	if err != nil {
		log.Fatalf("Can't create handler: %s", err)
	}
	defer contactHandler.Delete()

	contactHandler.Run(config.ServerPath)
}
