package valute

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code" validate:"required"`
	CharCode string  `xml:"CharCode" json:"char_code" validate:"required"`
	Value    float64 `xml:"Value" json:"value" validate:"required"`
}
