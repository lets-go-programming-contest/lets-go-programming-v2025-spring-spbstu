package currency

import "sort"

type Currency struct {
	NumCode  int     `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func SortCurrencies(currencies []Currency) []Currency {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
	return currencies;
}