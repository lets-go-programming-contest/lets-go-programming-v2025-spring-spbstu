package xmlToJson

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path/filepath"

	"github.com/go-playground/validator"
	"golang.org/x/net/html/charset"

	"github.com/kseniadobrovolskaia/task-3/internal/valute"
)

type Valutes struct {
	Vals []valute.Valute `xml:"Valute"`
}

func ReadXMLFile(inputFile string) ([]valute.Valute, error) {
	xmlFile, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	// A Decoder represents an XML parser reading a particular input stream. The parser assumes that its input is encoded in UTF-8.
	parser := xml.NewDecoder(xmlFile)
	// NewReaderLabel returns a reader that converts from the specified charset to UTF-8. The value found in the <?xml?> tag is passed to CharsetReader as the first argument (=label). The second argument is the reader itself. It uses Lookup to find the encoding that corresponds to label, and returns an error if Lookup returns nil. It is suitable for use as encoding/xml.Decoder's CharsetReader function.
	parser.CharsetReader = charset.NewReaderLabel

	var vals Valutes
	if err := parser.Decode(&vals); err != nil {
		return nil, errors.New("error in parse XML: " + err.Error())
	}

	// Check that all Valutes are valid
	for _, v := range vals.Vals {
		if err := validator.New().Struct(v); err != nil {
			return nil, errors.New(inputFile + ": validation failed due to: " + err.Error())
		}
	}
	return vals.Vals, nil
}

func WriteInJSONFile(outputFile string, vals []valute.Valute) (int, error) {
	if err := os.MkdirAll(filepath.Dir(outputFile), 0777); err != nil {
		return 0, err
	}

	jsonFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer jsonFile.Close()

	data, err := json.MarshalIndent(vals, "", " ")
	if err != nil {
		return 0, err
	}

	lenBytes, err := jsonFile.WriteString(string(data))
	if err != nil {
		return 0, err
	}
	return lenBytes, nil
}
