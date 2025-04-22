package encoder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"task-3/internal/data"
)

type JsonData struct {
	Data []Valute
}

type Valute struct {
	NumCode  string `json:"NumCode"`
	CharCode string `json:"CharCode"`
	Value    string `json:"Value"`
}

func Encode(elements data.DataElements, pathToOutput string) error {
	var jsonData JsonData
	for _, element := range elements {
		jsonData.Data = append(jsonData.Data, Valute{
			NumCode:  strconv.Itoa(element.NumCode),
			CharCode: element.CharCode,
			Value:    strconv.FormatFloat(element.Value, 'f', 1, 64),
		})
	}
	dir := filepath.Dir(pathToOutput)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory %v", err)
	}
	file, err := os.Create(pathToOutput)
	if err != nil {
		return fmt.Errorf("failed to create file %v", err)
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write JSON: %v", err)
	}
	return nil
}
