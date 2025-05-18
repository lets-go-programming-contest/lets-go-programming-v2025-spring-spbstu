package readconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Src string `yaml:"input-file"`
	Dst string `yaml:"output-file"`
}

func ReadConfig(configPath string) Config {
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
