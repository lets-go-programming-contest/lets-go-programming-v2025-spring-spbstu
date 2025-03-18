package xml2json

import (
	"encoding/xml"
	"strconv"
	"strings"
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

func ReadXML(data *xml.Decoder) []Format {
	var valCurs ValCurs
	err := data.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	var currencies []Format
	for _, v := range valCurs.Valute {
		value, err := strconv.ParseFloat(strings.Replace(v.Value, ",", ".", -1), 64)
		if err != nil {
			panic("Error during replacing comma")
		}

		num, err := strconv.Atoi(v.NumCode)
		if err != nil {
			panic("Error during atio")
		}

		currencies = append(currencies, Format{
			NumCode:  num,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return currencies
}
