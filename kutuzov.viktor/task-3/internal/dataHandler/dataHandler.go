package dataHandler

import (
	"log"
	"sort"
	"strconv"

	"github.com/vktr-ktzv/task3/internal/parser"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertAndSort(valCurs parser.ValCurs) []Currency {
	var currencies []Currency
	for _, valute := range valCurs.Valutes {

		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			log.Printf("Skipping Valute (invalid NumCode '%s'): %v", valute.NumCode, err)
			continue
		}

		value, err := valute.Value.ConvertToNormalFloat()
		if err != nil {
			log.Printf("Skipping Valute (invalid Value '%s'): %v", valute.Value, err)
			continue
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}
