package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int    `yaml:"serverPort"`
	Path string `yaml:"databasePath"`
}

func ParseConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read config file: %v", err)
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("fail to parse config file: %v", err)
	}
	if cfg.Port == 0 {
		return Config{}, fmt.Errorf("serverPort is not set")
	}
	if cfg.Path == "" {
		return Config{}, fmt.Errorf("databasePath is not set")
	}
	return cfg, nil
}
