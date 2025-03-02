package yaml_decoder

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfig struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Decode(filename string, out *YamlConfig) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("yaml_decoder.Decode %v: %w", filename, err)
	}
	defer file.Close()

	var YamlConfig YamlConfig

	if err = yaml.NewDecoder(file).Decode(&YamlConfig); err != nil {
		return fmt.Errorf("yaml_decoder.Decode %v: %w", filename, err)
	}

	if YamlConfig.InputFile == "" || YamlConfig.OutputFile == "" {
		return fmt.Errorf("yaml_decoder.Decode %v: Can't find `input-file` or `output-file`", filename)
	}

	*out = YamlConfig

	return nil
}
