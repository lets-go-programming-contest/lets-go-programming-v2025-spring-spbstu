package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/denisboborukhin/currency_converter/internal/currency"
)

func SaveCurrencies(outputFile string, currencies []currency.CurrencyJSON) error {
	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, data, 0644)
}
