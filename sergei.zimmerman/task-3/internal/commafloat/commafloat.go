package commafloat

import (
	"fmt"
	"strconv"
	"strings"
)

type CommaFloat float64

func (cf CommaFloat) ToString() string {
	return strings.ReplaceAll(fmt.Sprintf("%f", float64(cf)), ".", ",")
}

func (cf *CommaFloat) FromString(s string) error {
	normalized := strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("invalid number format: %s", s)
	}
	*cf = CommaFloat(f)
	return nil
}

func (cf *CommaFloat) UnmarshalText(text []byte) error {
	return (*cf).FromString(string(text))
}
