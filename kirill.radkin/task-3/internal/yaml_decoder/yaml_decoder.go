package yaml_decoder

import (
	"os"

	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	InputFile  string `yaml:"input-file" validate:"required"`
	OutputFile string `yaml:"output-file" validate:"required"`
}

func Decode(filename string) (*YamlConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var YamlConfig YamlConfig

	if err = yaml.NewDecoder(file).Decode(&YamlConfig); err != nil {
		return nil, err
	}

	if err = validator.New().Struct(YamlConfig); err != nil {
		return nil, err
	}

	return &YamlConfig, nil
}
