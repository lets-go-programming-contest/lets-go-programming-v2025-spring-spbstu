package currency

import "sort"

type Currency struct {
        NumCode  int     `xml:"NumCode" json:"num_code" validate:"required"`
        CharCode string  `xml:"CharCode" json:"char_code" validate:"required"`
        Value    string  `xml:"Value" json:"-" validate:"required"`
        FPValue  float64 `json:"value"`
}

type CurrencyList []Currency

func (l CurrencyList) Len() int {
        return len(l)
}

func (l CurrencyList) Less(idx1, idx2 int) bool {
        return l[idx1].FPValue < l[idx2].FPValue
}

func (l CurrencyList) Swap(idx1, idx2 int) {
        l[idx1], l[idx2] = l[idx2], l[idx1]
}

func (l CurrencyList) Sort() {
        sort.Sort(l)
}
