package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigRead  = errors.New("config read error")
	ErrConfigParse = errors.New("config parse error")
)

// Configuration structure for YAML parsing
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// LoadConfig reads and parses YAML configuration file
func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConfigRead, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrConfigParse, err)
	}

	return &cfg, nil
}
