package shortcurrency

import (
	"encoding/json"
	"os"
	"path"

	"task-3/internal/cbr"
)

type ShortCurrency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func FromCbr(c cbr.Currency) ShortCurrency {
	return ShortCurrency{c.NumCode, c.CharCode, float64(c.Value)}
}

const (
	modeRWXRXRX = 0755
	modeRWRR    = 0644
)

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, modeRWXRXRX)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func WriteJSON(filename string, currencies []cbr.Currency) error {
	shortsSlice := make([]ShortCurrency, len(currencies))

	for i, c := range currencies {
		shortsSlice[i] = FromCbr(c)
	}

	jsonData, err := json.MarshalIndent(shortsSlice, "", "  ")
	if err != nil {
		return err
	}

	err = ensureDir(path.Dir(filename))
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, modeRWRR)
	if err != nil {
		return err
	}

	return nil
}
