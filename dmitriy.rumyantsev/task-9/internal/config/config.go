package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `yaml:"server" validate:"required"`
	Database DatabaseConfig `yaml:"database" validate:"required"`
}

// ServerConfig contains server-related settings
type ServerConfig struct {
	Port    string        `yaml:"port" validate:"required"`
	Timeout time.Duration `yaml:"timeout" validate:"required,gt=0"`
}

// DatabaseConfig contains database connection settings
type DatabaseConfig struct {
	Host            string        `yaml:"host" validate:"required"`
	Port            string        `yaml:"port" validate:"required"`
	User            string        `yaml:"user" validate:"required"`
	Password        string        `yaml:"password"` // Password can be empty in some configurations
	Name            string        `yaml:"name" validate:"required"`
	MaxOpenConns    int           `yaml:"max_open_conns" validate:"required,gt=0"`
	MaxIdleConns    int           `yaml:"max_idle_conns" validate:"required,gt=0"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" validate:"required,gt=0"`
}

// Load loads configuration from the specified YAML file
func Load(path string) (*Config, error) {
	// Check if config file exists
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config file does not exist: %s", path)
	}

	// Read config file
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Parse YAML
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config file: %w", err)
	}

	// Initialize validator
	validate := validator.New()

	// Validate configuration using validator
	if err := validate.Struct(cfg); err != nil {
		// Format validation errors to be more user-friendly
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := formatValidationErrors(validationErrors)
		return nil, fmt.Errorf("invalid configuration: %s", strings.Join(errorMessages, "; "))
	}

	return &cfg, nil
}

// formatValidationErrors converts validator.ValidationErrors to more user-friendly messages
func formatValidationErrors(errors validator.ValidationErrors) []string {
	var errorMessages []string

	for _, err := range errors {
		var message string
		fieldName := formatFieldName(err.Namespace())

		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", fieldName)
		case "gt":
			message = fmt.Sprintf("%s must be greater than %s", fieldName, err.Param())
		case "oneof":
			message = fmt.Sprintf("%s must be one of: %s", fieldName, err.Param())
		default:
			message = fmt.Sprintf("%s failed validation: %s", fieldName, err.Tag())
		}

		errorMessages = append(errorMessages, message)
	}

	return errorMessages
}

// formatFieldName formats field path (e.g. Config.Server.Port) to a more readable form (e.g. server.port)
func formatFieldName(fieldPath string) string {
	// Remove the top-level struct name
	parts := strings.Split(fieldPath, ".")
	if len(parts) > 1 {
		parts = parts[1:]
	}

	// Convert to lowercase
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}

	return strings.Join(parts, ".")
}
