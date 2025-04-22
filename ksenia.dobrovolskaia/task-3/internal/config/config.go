package config

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-playground/validator"
)

type Config struct {
	InputFile  string `json:"input-file" validate:"required"`
	OutputFile string `json:"output-file" validate:"required"`
}

func ReadConfigFile(configPath string) (*Config, error) {
	//fmt.Printf("config file: %s\n", *configPath)
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			return nil, errors.New(configPath + "syntax error at byte offset" + string(e.Offset))
		}
		panic(err)
	}
	err = validator.New().Struct(config)
	if err != nil {
		return nil, errors.New(configPath + ": validation failed due to: " + err.Error())
	}
	//fmt.Printf("decoded config: %+v\n", config)
	return &config, nil
}
