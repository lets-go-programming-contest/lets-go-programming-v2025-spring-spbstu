package configLoader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file" validate:"required"`
	OutputFile string `yaml:"output-file" validate:"required"`
}

func LoadConfig(path string) Config {
	configFile, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading config: %v", err))
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		panic(fmt.Sprintf("Error parsing YAML: %v", err))
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("Config must contain input-file and output-file")
	}

	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic("Input file does not exist")
	}

	return config
}
