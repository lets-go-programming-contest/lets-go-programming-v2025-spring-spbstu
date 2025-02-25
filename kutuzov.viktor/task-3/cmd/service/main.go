package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.Parse()

	// Загрузка конфигурации
	config := loadConfig(configPath)

	// Парсинг XML
	valCurs := parseXML(config.InputFile)

	// Конвертация и сортировка
	currencies := convertAndSort(valCurs)

	// Сохранение результатов
	saveJSON(config.OutputFile, currencies)
}

func loadConfig(path string) Config {
	configFile, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading config: %v", err))
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		panic(fmt.Sprintf("Error parsing YAML: %v", err))
	}

	if config.InputFile == "" || config.OutputFile == "" {
		panic("Config must contain input-file and output-file")
	}

	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic("Input file does not exist")
	}

	return config
}

func parseXML(inputFile string) ValCurs {
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Sprintf("Error opening XML: %v", err))
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	var valCurs ValCurs
	if err := xml.Unmarshal(byteValue, &valCurs); err != nil {
		panic(fmt.Sprintf("Error parsing XML: %v", err))
	}

	return valCurs
}

func convertAndSort(valCurs ValCurs) []Currency {
	var currencies []Currency
	for _, valute := range valCurs.Valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			panic(fmt.Sprintf("Invalid NumCode: %v", err))
		}

		valueStr := strings.Replace(valute.Value, ",", ".", 1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(fmt.Sprintf("Invalid Value: %v", err))
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}

func saveJSON(outputFile string, currencies []Currency) {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Error creating dir: %v", err))
	}

	file, err := os.Create(outputFile)
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(currencies); err != nil {
		panic(fmt.Sprintf("Error encoding JSON: %v", err))
	}
}
