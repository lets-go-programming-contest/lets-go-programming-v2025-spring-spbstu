package cursesparser

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"os"
	"strconv"
	"strings"
)

type FloatString float64

func (f *FloatString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	s = strings.Replace(s, ",", ".", -1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f = FloatString(val)
	return nil
}

type valCurs struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Curs    []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    			`xml:"NumCode"`
	CharCode string 			`xml:"CharCode"`
	Value    FloatString 	`xml:"Value"`
}

func Parse(cursesFile *os.File) ([]Valute, error) {
	decoder := xml.NewDecoder(cursesFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var curses valCurs
	err := decoder.Decode(&curses)
	if err != nil {
		return nil, err
	}

	return curses.Curs, nil
}
