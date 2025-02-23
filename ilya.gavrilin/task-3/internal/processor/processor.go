package processor

import (
	"sort"
	"strconv"

	"task-3/internal/fetcher"
)

// CurrencyOutput represents the final output structure.
type CurrencyOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

// SortCurrencies sorts currency data in descending order of value.
func SortCurrencies(currencies []fetcher.Currency) []CurrencyOutput {
	var output []CurrencyOutput
	for _, c := range currencies {
		numCode, err := strconv.Atoi(c.NumCode)
		if err != nil {
			continue // Ignore invalid entries
		}
		output = append(output, CurrencyOutput{
			NumCode:  numCode,
			CharCode: c.CharCode,
			Value:    c.Value,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})

	return output
}