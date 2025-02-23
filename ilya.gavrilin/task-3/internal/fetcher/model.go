package fetcher

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Currency represents a single currency entry from the XML file.
type Currency struct {
	NumCode  string  `xml:"NumCode" validate:"required,numeric"`
	CharCode string  `xml:"CharCode" validate:"required,alpha,len=3"`
	Value    float64 `validate:"required,gt=0"`
}

// Custom UnmarshalXML to fix comma issue in float values.
func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var temp struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	if err := d.DecodeElement(&temp, &start); err != nil {
		return err
	}

	c.NumCode = temp.NumCode
	c.CharCode = temp.CharCode

	// Replace "," with "." before converting to float
	value := strings.Replace(temp.Value, ",", ".", 1)
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value %s: %w", temp.Value, err)
	}

	c.Value = parsedValue
	return nil
}

// ValCurs represents the root XML structure.
type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}

// validateCurrency checks if currency data is valid.
func validateCurrency(c Currency) error {
	validate := validator.New()
	return validate.Struct(c)
}