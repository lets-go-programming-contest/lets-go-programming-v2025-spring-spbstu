package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3" // Сторонний пакет для YAML
)

// Структура конфигурации
type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

// Загрузка конфига из YAML-файла
func LoadConfig(configPath string) Config {
	// Чтение файла
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка чтения конфига: %v", err))
	}

	// Парсинг YAML
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("Ошибка парсинга YAML: %v", err))
	}

	return cfg
}
