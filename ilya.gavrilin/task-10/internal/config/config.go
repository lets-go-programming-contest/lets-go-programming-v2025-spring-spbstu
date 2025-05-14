package config

import (
	"os"
	"strconv"
)

type Config struct {
	GRPC GRPCConfig
	REST RESTConfig
	DB   DBConfig
}

type GRPCConfig struct {
	Port string
}

type RESTConfig struct {
	Port string
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		dbPort = 5432
	}

	return &Config{
		GRPC: GRPCConfig{
			Port: getEnv("GRPC_PORT", "50051"),
		},
		REST: RESTConfig{
			Port: getEnv("REST_PORT", "8080"),
		},
		DB: DBConfig{
			Driver:   getEnv("DB_DRIVER", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "contacts"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
