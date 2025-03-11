package config

import (
	"fmt"
	"os"

	"github.com/dmitriy.rumyantsev/task-3/internal/models"
	my_validator "github.com/dmitriy.rumyantsev/task-3/internal/validator"

	validator "github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Reads and parses YAML configuration
func ReadConfig(path string, validate *validator.Validate) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML configuration: %w", err)
	}

	err = validate.Struct(config)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return nil, fmt.Errorf("configuration validation errors: %s",
				my_validator.FormatValidationErrors(validationErrors))
		}
		return nil, fmt.Errorf("configuration validation error: %w", err)
	}

	return &config, nil
}
