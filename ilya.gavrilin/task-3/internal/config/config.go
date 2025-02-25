package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration structure.
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// LoadConfig reads and parses the YAML configuration file.
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Validate required fields
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return nil, fmt.Errorf("missing required fields: input-file or output-file")
	}

	return &cfg, nil
}