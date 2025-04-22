package models

import (
	"encoding/xml"
)

// Structures for XML parsing
type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	ID       string   `xml:"ID,attr"`
	NumCode  int      `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  int      `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
}

// YAML Configuration structure
type Config struct {
	InputFile  string `yaml:"input-file" validate:"required,file"`
	OutputFile string `yaml:"output-file" validate:"required"`
}

// JSON Output structure
type OutputValute struct {
	NumCode  int     `json:"num_code" validate:"required,gt=0"`
	CharCode string  `json:"char_code" validate:"required,len=3"`
	Value    float64 `json:"value" validate:"required,gt=0"`
}
