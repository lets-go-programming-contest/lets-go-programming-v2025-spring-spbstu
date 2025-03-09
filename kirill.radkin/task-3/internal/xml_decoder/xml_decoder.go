package xmldecoder

import (
	"encoding/xml"
	"os"

	"github.com/yanelox/task-3/internal/mytypes"
	"golang.org/x/net/html/charset"
	"gopkg.in/go-playground/validator.v9"
)

func Decode(filename string) (*mytypes.ValCurs, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ValCurs mytypes.ValCurs

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err = decoder.Decode(&ValCurs); err != nil {
		return nil, err
	}

	validate := validator.New()

	if err = validate.Struct(ValCurs); err != nil {
		return nil, err
	}

	for _, valute := range ValCurs.Valutes {
		if err = validate.Struct(valute); err != nil {
			return nil, err
		}
	}

	return &ValCurs, nil
}
