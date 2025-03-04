package processor

import (
	"sort"

	"task-3/internal/fetcher"
)

// SortCurrencies sorts currency data in descending order of value.
func SortCurrencies(currencies []fetcher.Currency) []fetcher.Currency {
	sort.Slice(currencies, func(i, j int) bool {
		// Directly compare CurrencyValue values (their underlying type is float64).
		return currencies[i].Value > currencies[j].Value
	})
	return currencies
}
