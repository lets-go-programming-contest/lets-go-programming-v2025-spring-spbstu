package xmldecoder

import (
	"encoding/xml"
	"fmt"
	"os"

	vl "github.com/yanelox/task-3/internal/valute"
	"golang.org/x/net/html/charset"
)

type xmlValute struct {
	ID        vl.ID         `xml:"ID,attr"`
	NumCode   vl.NumCode    `xml:"NumCode"`
	CharCode  vl.CharCode   `xml:"CharCode"`
	Nominal   vl.Nominal    `xml:"Nominal"`
	Name      vl.ValuteName `xml:"Name"`
	Value     vl.Value      `xml:"Value"`
	VunitRate vl.VunitRate  `xml:"VunitRate"`
}

type xmlValCurs struct {
	XMLName xml.Name       `xml:"ValCurs"`
	Date    vl.Date        `xml:"Date,attr"`
	Name    vl.ValCursName `xml:"name,attr"`
	Valutes []xmlValute    `xml:"Valute"`
}

func Decode(filename string, out *vl.ValCurs) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("xmldecoder.Decode %v: %w", filename, err)
	}
	defer file.Close()

	var xmlValCurs xmlValCurs

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	if err = decoder.Decode(&xmlValCurs); err != nil {
		return fmt.Errorf("xmldecoder.Decode %v: %w", filename, err)
	}

	if xmlValCurs.Date != "" {
		out.Date = xmlValCurs.Date
	} else {
		return fmt.Errorf("xmldecoder.Decode %v: empty ValCurs Date\n", filename)
	}

	if xmlValCurs.Name != "" {
		out.Name = xmlValCurs.Name
	} else {
		return fmt.Errorf("xmldecoder.Decode %v: empty ValCurs Name\n", filename)
	}

	for i, v := range xmlValCurs.Valutes {
		if v.ID == "" || v.NumCode == "" || v.CharCode == "" || v.Nominal == "" || v.Name == "" || v.Value == "" || v.VunitRate == "" {
			return fmt.Errorf("xmldecoder.Decode %v: incomplete Valute %d\n", filename, i)
		}

		out.Valutes = append(out.Valutes, vl.Valute{v.ID, v.NumCode, v.CharCode, v.Nominal, v.Name, v.Value, v.VunitRate})
	}

	return nil
}
