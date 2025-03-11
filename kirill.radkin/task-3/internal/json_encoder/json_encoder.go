package jsonencoder

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/yanelox/task-3/internal/mytypes"
	"gopkg.in/go-playground/validator.v9"
)



func Encode(filename string, valutes []mytypes.Valute) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	vaidate := validator.New()
	for _, valute := range valutes {
		if err = vaidate.Struct(valute); err != nil {
			return err
		}
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err = encoder.Encode(&valutes); err != nil {
		return err
	}

	return nil
}
