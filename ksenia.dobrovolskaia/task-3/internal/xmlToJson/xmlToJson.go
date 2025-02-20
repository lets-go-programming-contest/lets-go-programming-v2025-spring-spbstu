package xmlToJson

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path/filepath"

	"github.com/go-playground/validator"
	"golang.org/x/net/html/charset"

	"github.com/kseniadobrovolskaia/task-3/internal/valute"
)

func ReadXMLFile(inputFile string) ([]valute.Valute, error) {
	vals := make([]valute.Valute, 0)

	xmlFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	parser := xml.NewDecoder(xmlFile)
	parser.CharsetReader = charset.NewReaderLabel

	for {
		t, _ := parser.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "Valute" {
				var item valute.Valute
				parser.DecodeElement(&item, &se)
				err = validator.New().Struct(item)
				if err != nil {
					return nil, errors.New(inputFile + ": validation failed due to: " + err.Error())
				}
				//fmt.Printf("Add valute with value: %f\n", item.Value)
				vals = append(vals, item)
			}
		}
	}
	//fmt.Printf("decoded valutes: %+v\n", vals)
	return vals, nil
}

func WriteInJSONFile(outputFile string, vals []valute.Valute) (int, error) {
	if err := os.MkdirAll(filepath.Dir(outputFile), 0777); err != nil {
		return 0, err
	}

	jsonFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer jsonFile.Close()

	data, err := json.MarshalIndent(vals, "", " ")
	if err != nil {
		return 0, err
	}

	lenBytes, err := jsonFile.WriteString(string(data))
	if err != nil {
		return 0, err
	}
	return lenBytes, nil
}
