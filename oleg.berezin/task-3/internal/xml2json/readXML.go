package xml2json

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func readXML(data []byte) []Format {
	var valCurs ValCurs
	err := xml.Unmarshal(data, &valCurs)
	if err != nil {
		panic("Error during unmarshal xml")
	}

	var currencies []Format
	for _, v := range valCurs.Valutes {
		value, err := strconv.ParseFloat(strings.Replace(v.Value, ",", ".", -1), 64)
		if err != nil {
			panic("Error during replacing comma")
		}

		currencies = append(currencies, Format{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return currencies
}
