package cbr

import (
	"bytes"
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	Date   string `xml:"Date,attr"`
	Name   string `xml:"name,attr"`
	Valute []struct {
		ID        string       `xml:"ID,attr"`
		NumCode   int          `xml:"NumCode"`
		CharCode  string       `xml:"CharCode"`
		Nominal   string       `xml:"Nominal"`
		Name      string       `xml:"Name"`
		Value     russianFloat `xml:"Value"`
		VunitRate string       `xml:"VunitRate"`
	} `xml:"Valute"`
}

type russianFloat float64

func normalizeAmerican(old string) string {
	return strings.ReplaceAll(old, ",", ".")
}

func (val *russianFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valString string
	err := d.DecodeElement(&valString, &start)
	if err != nil {
		return err
	}
	valString = normalizeAmerican(valString)
	parsedVal, err := strconv.ParseFloat(valString, 64)
	if err != nil {
		return err
	}
	*val = russianFloat(parsedVal)
	return nil
}

func ReadCbrXml(xmlPath string) (*ValCurs, error) {
	data, err := os.ReadFile(xmlPath)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader([]byte(data))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	valCurs := new(ValCurs)
	err = d.Decode(&valCurs)
	if err != nil {
		return nil, err
	}

	return valCurs, nil
}
