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
	NumCode  int64   `xml:"NumCode" validate:"required,numeric"`
	CharCode string  `xml:"CharCode" validate:"required,alpha,len=3"`
	Value    float64 `validate:"required,gt=0"`
}

// Custom UnmarshalXML to fix comma issue in float values.
func (c *Currency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Currency
	aux := &struct {
		Alias
		Value string `xml:"Value"`
	}{Alias: Alias(*c)} // Embed existing data to avoid overwriting

	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	// Replace "," with "." and parse as float
	value := strings.Replace(aux.Value, ",", ".", 1)
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value %s: %w", aux.Value, err)
	}

	// Assign parsed data back to the original struct
	*c = Currency(aux.Alias)
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
