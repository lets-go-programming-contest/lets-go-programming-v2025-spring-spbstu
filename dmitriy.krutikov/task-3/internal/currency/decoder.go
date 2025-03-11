package currency

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

func Decode(filePath string) ([]Currency, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	
	reader := bytes.NewReader(file)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, err
	}
	var currencies []Currency
	for _, valute := range valCurs.Valutes {
		valueStr := strings.Replace(valute.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, err
		}

		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})

		
	}
	return currencies, nil
}

func PrintCurrencies(currencies []Currency) {
	for _, currency := range currencies {
		fmt.Printf("NumCode: %d, CharCode: %s, Value: %.4f\n", currency.NumCode, currency.CharCode, currency.Value)
	}
}