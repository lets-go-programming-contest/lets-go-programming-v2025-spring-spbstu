package configparser

import (
	"gopkg.in/yaml.v2"
	"errors"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Parse(data []byte) (Config, error) {
	var config Config

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, errors.Join(errors.New("Failed to unmarshal config"), err)
	}

	return config, nil
}
