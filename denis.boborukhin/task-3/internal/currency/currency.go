package currency

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type CurrencyJSON struct {
	NumCode  string  `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type Currencies struct {
	Currencies []CurrencyJSON `xml:"Valute"`
}

func (c *Currencies) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type tmpCurrency struct {
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		ValueStr string `xml:"Value"`
	}

	token, err := d.Token()
	for ; err == nil; token, err = d.Token() {
		if ee, ok := token.(xml.EndElement); ok && ee.Name.Local == start.Name.Local {
			return nil
		}

		se, ok := token.(xml.StartElement)
		if !ok || se.Name.Local != "Valute" {
			continue
		}

		if ee, ok := token.(xml.EndElement); ok && ee.Name.Local == start.Name.Local {
			break
		}

		// if begin of the <Valute>
		var tmp tmpCurrency
		if err := d.DecodeElement(&tmp, &se); err != nil {
			return err
		}

		if tmp.NumCode == "" || tmp.CharCode == "" || tmp.ValueStr == "" {
			continue // skip
		}

		valueStr := strings.Replace(tmp.ValueStr, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			continue // skip
		}

		c.Currencies = append(c.Currencies, CurrencyJSON{
			NumCode:  tmp.NumCode,
			CharCode: tmp.CharCode,
			Value:    value,
		})
	}

	if err != io.EOF {
		return err
	}

	return nil
}

func ConvertToJSON(inputFile string) ([]CurrencyJSON, error) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}

	var currencies Currencies
	if err := decoder.Decode(&currencies); err != nil {
		return nil, err
	}

	return currencies.Currencies, nil
}

func SortCurrenciesByValue(currencies []CurrencyJSON) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
