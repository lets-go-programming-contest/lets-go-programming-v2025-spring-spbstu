package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"ID,attr"`
		NumCode   string `xml:"NumCode"`
		CharCode  string `xml:"CharCode"`
		Nominal   string `xml:"Nominal"`
		Name      string `xml:"Name"`
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
	} `xml:"Valute"`
}

func main() {
	filePath := "example_data/1.xml"

	// Read the file as a byte string
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	r := bytes.NewReader([]byte(content))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel

	valCurs := new(ValCurs)
	err = d.Decode(&valCurs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(valCurs)
}
