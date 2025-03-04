package fetcher

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// CurrencyValue is a custom type that unmarshals text with a comma as decimal separator.
type CurrencyValue float64

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (cv *CurrencyValue) UnmarshalText(text []byte) error {
	// Replace comma with dot.
	s := strings.Replace(string(text), ",", ".", 1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value %s: %w", text, err)
	}
	*cv = CurrencyValue(f)
	return nil
}

// Currency represents a single currency entry from the XML file.
type Currency struct {
	NumCode  int64         `xml:"NumCode" validate:"required,numeric"`
	CharCode string        `xml:"CharCode" validate:"required,alpha,len=3"`
	Value    CurrencyValue `xml:"Value" validate:"required,gt=0"`
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
