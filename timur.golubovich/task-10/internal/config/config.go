package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Host     string `yaml:"DB_HOST"`
	User     string `yaml:"DB_USER"`
	Password string `yaml:"DB_PASSWORD"`
	Name     string `yaml:"DB_NAME"`
	DBPort   int    `yaml:"DB_PORT"`
	GRPCPort int    `yaml:"GRPC_PORT"`
	HTTPPort int    `yaml:"HTTP_PORT"`
}

func ParseConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read config file: %v", err)
	}
	cfg := Config{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Name:     "phonebook",
		DBPort:   5432,
		GRPCPort: 50051,
		HTTPPort: 8080,
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("fail to parse config file: %v", err)
	}
	return cfg, nil
}

func MakeString(cfg Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.DBPort)
}
