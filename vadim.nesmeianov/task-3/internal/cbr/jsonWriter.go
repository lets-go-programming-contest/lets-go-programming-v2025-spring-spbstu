package cbr

import (
	"encoding/json"
	"sort"
)

type ValCursShort struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type JsonArr []ValCursShort

func (a JsonArr) Len() int           { return len(a) }
func (a JsonArr) Less(i, j int) bool { return a[i].Value > a[j].Value }
func (a JsonArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func converValCursToJsonArr(oldData *ValCurs) *JsonArr {
	n := len(oldData.Valute)
	newData := make(JsonArr, 0, n)
	for i := range n {
		newData = append(newData, ValCursShort{
			NumCode:  oldData.Valute[i].NumCode,
			CharCode: oldData.Valute[i].CharCode,
			Value:    float64(oldData.Valute[i].Value),
		})
	}

	return &newData
}

func GetSortedJsonString(oldData *ValCurs) (*[]byte, error) {
	valCursJson := converValCursToJsonArr(oldData)
	sort.Sort(valCursJson)

	jsonData, err := json.MarshalIndent(valCursJson, "", "\t")
	if err != nil {
		return nil, err
	}

	return &jsonData, nil
}
