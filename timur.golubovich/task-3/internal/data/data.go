package data

type RawDataElement struct {
	NumCode  string
	CharCode string
	Value    string
}

type RawDataElements []RawDataElement

type DataElement struct {
	NumCode  int
	CharCode string
	Value    float64
}

type DataElements []DataElement
