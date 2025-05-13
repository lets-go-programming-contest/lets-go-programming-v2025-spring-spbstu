package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DbHost  string `json:"DbHost"`
	DbPort  string `json:"DbPort"`
	DbUser  string `json:"DbUser"`
	DbPswrd string `json:"DbPswrd"`
	DbName  string `json:"DbName"`
	Port    string `json:"Port"`
}

func ReadConfigFile(configPath string) (*Config, error) {
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
	return &config, nil
}
