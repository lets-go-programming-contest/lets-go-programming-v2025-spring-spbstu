package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigRead  = errors.New("config read error")
	ErrConfigParse = errors.New("config parse error")
)

const configDir = "internal/config" // Directory for configs
const configFile = "config.yaml"    // Configuration file name

// Configuration structure for YAML parsing
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// LoadConfig reads and parses YAML configuration file
func LoadConfig() (*Config, error) {
	configPath := filepath.Join(configDir, configFile)

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
