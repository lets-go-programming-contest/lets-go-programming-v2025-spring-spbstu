package decoder

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"task-3/internal/data"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"ID,attr"`
		NumCode   string `xml:"NumCode"`
		CharCode  string `xml:"CharCode"`
		Nominal   string `xml:"Nominal"`
		Name      string `xml:"Name"`
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
	} `xml:"Valute"`
}

func Decode(pathToInput string) (data.RawDataElements, error) {
	content, err := os.ReadFile(pathToInput)
	if err != nil {
		return data.RawDataElements{}, err
	}
	reader := bytes.NewReader([]byte(content))
	xmlDecoder := xml.NewDecoder(reader)
	xmlDecoder.CharsetReader = charset.NewReaderLabel
	var valCurs ValCurs
	err = xmlDecoder.Decode(&valCurs)
	if err != nil {
		return data.RawDataElements{}, fmt.Errorf("fail to decode: %v", err)
	}
	var rawElements data.RawDataElements
	for _, valute := range valCurs.Valute {
		rawElements = append(rawElements, data.RawDataElement{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		})
	}
	return rawElements, nil
}
