package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type FloatFromRussia string

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string          `xml:"NumCode"`
	CharCode string          `xml:"CharCode"`
	Value    FloatFromRussia `xml:"Value"`
}

func (value *FloatFromRussia) ConvertToNormalFloat() (float64, error) {
	valueStr := strings.Replace(string(*value), ",", ".", 1)
	parsedValue, err := strconv.ParseFloat(valueStr, 64)

	return parsedValue, err
}

func ParseXML(inputFile string) (ValCurs, error) {
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Sprintf("Error opening XML: %v", err))
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel //Set encoding handler
	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return ValCurs{}, fmt.Errorf("failed to parse XML: %w", err)
	}

	return valCurs, nil
}
