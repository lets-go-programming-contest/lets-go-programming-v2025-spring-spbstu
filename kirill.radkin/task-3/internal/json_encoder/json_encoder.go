package jsonencoder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	vl "github.com/yanelox/task-3/internal/valute"
)

type jsonValute struct {
	NumCode  vl.NumCode  `json:"num_code"`
	CharCode vl.CharCode `json:"char_code"`
	Value    vl.Value    `json:"value"`
}

func Encode(filename string, valutes []vl.Valute) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("jsonencoder.Encode %v: %w", filename, err)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return fmt.Errorf("jsonencoder.Encode %v: %w", filename, err)
	}
	defer file.Close()

	var jsonValutes []jsonValute
	for _, v := range valutes {
		if v.NumCode == "" || v.CharCode == "" || v.Value == "" {
			return fmt.Errorf("jsonencoder.Encode %v: incomplete Valute", filename)
		}

		jsonValutes = append(jsonValutes, jsonValute{v.NumCode, v.CharCode, v.Value})
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err = encoder.Encode(&jsonValutes); err != nil {
		return fmt.Errorf("jsonencoder.Encode %v: %w", filename, err)
	}

	return nil
}
