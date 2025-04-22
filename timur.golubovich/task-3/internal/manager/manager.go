package manager

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func ParseConfig(pathToConfig string) (Config, error) {
	configFile, err := os.ReadFile(pathToConfig)
	if err != nil {
		return Config{}, fmt.Errorf("fail to read config file: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("fail to parse config file: %v", err)
	}
	return config, nil
}
