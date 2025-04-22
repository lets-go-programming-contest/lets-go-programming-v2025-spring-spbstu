package xmlparser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"task-3/internal/currency"

	"golang.org/x/text/encoding/charmap"
)

var (
	ErrFileOpen  = errors.New("file open error")
	ErrXMLDecode = errors.New("xml decode error")
)

// Custom type to handle currency value parsing
type CurrencyValue float64

// Implement TextUnmarshaler interface for custom parsing
func (cv *CurrencyValue) UnmarshalText(text []byte) error {
	s := strings.Replace(string(text), ",", ".", 1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("invalid currency value: %w", err)
	}
	*cv = CurrencyValue(f)
	return nil
}

// Temporary XML structure with custom value type
type xmlCurrency struct {
	NumCode  int           `xml:"NumCode"`
	CharCode string        `xml:"CharCode"`
	Value    CurrencyValue `xml:"Value"`
}

// XML root element structure
type valCurs struct {
	Currencies []xmlCurrency `xml:"Valute"`
}

func ParseXML(filePath string) ([]currency.Currency, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFileOpen, err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	var vc valCurs
	if err := decoder.Decode(&vc); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrXMLDecode, err)
	}

	// Convert to final currency structure
	result := make([]currency.Currency, len(vc.Currencies))
	for i, c := range vc.Currencies {
		result[i] = currency.Currency{
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Value:    float64(c.Value), // Convert custom type to float64
		}
	}

	return result, nil
}
