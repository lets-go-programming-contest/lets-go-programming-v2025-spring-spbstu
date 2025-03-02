package valute

type ID string
type NumCode string
type CharCode string
type Nominal string
type ValuteName string
type Value string
type VunitRate string

type Valute struct {
	ID        ID
	NumCode   NumCode
	CharCode  CharCode
	Nominal   Nominal
	Name      ValuteName
	Value     Value
	VunitRate VunitRate
}

type Date string
type ValCursName string

type ValCurs struct {
	Date    Date
	Name    ValCursName
	Valutes []Valute
}
