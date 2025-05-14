package currency

import "sort"

// Currency structure with XML/JSON tags
type Currency struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

// SortCurrencies sorts currencies in descending order by Value
func SortCurrencies(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
