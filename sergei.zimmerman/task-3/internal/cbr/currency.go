package cbr

import (
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
	"github.com/go-playground/validator/v10"

	"task-3/internal/commafloat"
)

type CurrencyValues struct {
	Date       string     `xml:"Date,attr" validate:"required"`
	Name       string     `xml:"name,attr" validate:"required"`
	Currencies []Currency `xml:"Valute" validate:"dive"`
}

type Currency struct {
	Id        string                `xml:"ID,attr" validate:"required"`
	NumCode   int                   `xml:"NumCode"`
	CharCode  string                `xml:"CharCode"`
	Nominal   int                   `xml:"Nominal"`
	Name      string                `xml:"Name"`
	Value     commafloat.CommaFloat `xml:"Value"`
	VunitRate commafloat.CommaFloat `xml:"VunitRate"`
}

func ParseCbrXML(filename string) (*CurrencyValues, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := xml.NewDecoder(file)

	/* Handle whatever nonsense encoding windows has.
	   https://pkg.go.dev/golang.org/x/net/html/charset
	   > NewReaderLabel returns a reader that converts from the specified
	   > charset to UTF-8. It uses Lookup to find the encoding that corresponds to
	   > label, and returns an error if Lookup returns nil. It is suitable for use as
	   > encoding/xml.Decoder's CharsetReader function. */

	decoder.CharsetReader = charset.NewReaderLabel
	valCurs := new(CurrencyValues)
	err = decoder.Decode(&valCurs)

	if err != nil {
		return nil, err
	}

	validator := validator.New()
	err = validator.Struct(valCurs)

	if err != nil {
		return nil, err
	}

	return valCurs, nil
}
