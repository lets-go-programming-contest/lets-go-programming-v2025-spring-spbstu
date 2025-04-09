package xml2json

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
)

type Value string

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    Value  `xml:"Value"`
	} `xml:"Valute"`
}

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rawValue string
	if err := d.DecodeElement(&rawValue, &start); err != nil {
		return err
	}
	*v = Value(strings.ReplaceAll(rawValue, ",", "."))
	return nil
}

func ReadXML(data *xml.Decoder) ([]Format, error) {
	var valCurs ValCurs
	err := data.Decode(&valCurs)
	if err != nil {
		return []Format{}, errors.New("error during decoding xml")
	}

	var currencies []Format
	for _, v := range valCurs.Valute {
		value, err := strconv.ParseFloat(string(v.Value), 64)
		if err != nil {
			return []Format{}, errors.New("error during replacing comma")
		}

		num, err := strconv.Atoi(v.NumCode)
		if err != nil {
			return []Format{}, errors.New("error during atio")
		}

		currencies = append(currencies, Format{
			NumCode:  num,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}
