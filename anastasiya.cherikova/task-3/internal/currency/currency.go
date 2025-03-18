package currency

import "sort"

// Структура валюты (теги для XML и JSON)
type Currency struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

// Сортировка валют по убыванию значения
func SortCurrencies(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
