package mytypes

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type MyFloat64 float64

type Valute struct {
	NumCode   string    `xml:"NumCode" json:"num_code" validate:"required"`
	CharCode  string    `xml:"CharCode" json:"char_code" validate:"required"`
	Value     MyFloat64 `xml:"Value" json:"value" validate:"required"`
}

func (val *MyFloat64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	s = strings.Replace(s, ",", ".", 1)
	new_val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*val = MyFloat64(new_val)

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs" validate:"required"`
	Valutes []Valute `xml:"Valute"  validate:"required"`
}
