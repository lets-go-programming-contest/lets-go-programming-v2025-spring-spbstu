package fetcher

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html/charset"
)

// FetchData loads and parses XML data.
func FetchData(data io.Reader) ([]Currency, error) {
	// Create XML decoder with charset support
	decoder := xml.NewDecoder(data)
	decoder.CharsetReader = charset.NewReaderLabel // Fix windows-1251 encoding issue

	// Parse XML
	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	// Validate and filter valid records
	var validCurrencies []Currency
	for _, c := range valCurs.Currencies {
		if err := validateCurrency(c); err == nil {
			validCurrencies = append(validCurrencies, c)
		}
	}

	if len(validCurrencies) == 0 {
		return nil, errors.New("no valid currency data found")
	}

	return validCurrencies, nil
}