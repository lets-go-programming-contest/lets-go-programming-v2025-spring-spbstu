package cursesparser

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"os"
	"strconv"
	"strings"
)

type valCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Curs    []xmlValute `xml:"Valute"`
}

type xmlValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Valute struct {
	NumCode  int
	CharCode string
	Value    float64
}

func xmlValute2Valute(valute xmlValute) (Valute, error) {
	valueStr := strings.Replace(valute.Value, ",", ".", -1)

	valueFloat, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return Valute{0, "", 0}, err
	}

	return Valute{valute.NumCode, valute.CharCode, valueFloat}, nil
}

func Parse(cursesFile *os.File) ([]Valute, error) {
	var curses valCurs

	decoder := xml.NewDecoder(cursesFile)
	decoder.CharsetReader = charset.NewReaderLabel

	err := decoder.Decode(&curses)
	if err != nil {
		return []Valute{}, err
	}

	var valutes []Valute
	for _, xmlvalute := range curses.Curs {
		valute, err := xmlValute2Valute(xmlvalute)
		if err != nil {
			return []Valute{}, err
		}

		valutes = append(valutes, valute)
	}

	return valutes, nil
}
