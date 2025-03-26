package currency

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type comaSepFloat64 float64

var errValUnmarshFailed = errors.New("failed unmarshalling a 'Value' float")

// this method is implicitly used by `decoder.DecodeElement()`
func (f *comaSepFloat64) UnmarshalText(text []byte) error {
        textComaToDot := strings.ReplaceAll(string(text), `,`, `.`)

        parsedVal, err := strconv.ParseFloat(textComaToDot, 64)
        if err != nil {
                return errors.Join(errValUnmarshFailed, err)
        }

        *f = comaSepFloat64(parsedVal)
        return nil
}

type CurrencyDat struct {
        NumCode  int            `xml:"NumCode" json:"num_code" validate:"required"`
        CharCode string         `xml:"CharCode" json:"char_code" validate:"required"`
        Value    comaSepFloat64 `xml:"Value" json:"value" validate:"required"`
}

type Currency struct {
        Curr CurrencyDat `xml:"Valute"`
}

type CurrencyList []Currency

type Scheme struct {
        List CurrencyList `xml:"ValCurs"`
}

func NewScheme() Scheme {
        return Scheme{
                List: make(CurrencyList, 0),
        }
}

func (l CurrencyList) Len() int {
        return len(l)
}

func (l CurrencyList) Less(idx1, idx2 int) bool {
        return l[idx1].Curr.Value < l[idx2].Curr.Value
}

func (l CurrencyList) Swap(idx1, idx2 int) {
        l[idx1], l[idx2] = l[idx2], l[idx1]
}

func (l CurrencyList) Sort() {
        sort.Sort(l)
}
