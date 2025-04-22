package valute

import (
	"fmt"
	"strconv"
	"strings"
)

type FloatRuss float64

type Valute struct {
	NumCode  int       `xml:"NumCode" json:"num_code" validate:"required"`
	CharCode string    `xml:"CharCode" json:"char_code" validate:"required"`
	Value    FloatRuss `xml:"Value" json:"value" validate:"required"`
}

// Implement encoding.TextUnmarshaler interface
func (fr *FloatRuss) UnmarshalText(data []byte) error {
	value := strings.Replace(string(data), ",", ".", 1)
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("error in parse float %s: %w", value, err)
	}
	*fr = FloatRuss(f)
	return nil
}
