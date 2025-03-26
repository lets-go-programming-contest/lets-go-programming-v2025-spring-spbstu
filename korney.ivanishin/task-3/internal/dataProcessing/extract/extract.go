package extract

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/currency"
	"golang.org/x/net/html/charset"
)

var (
        errOpenInpFileFailed   = errors.New("failed opening input file")
        errDecodeInpFileFailed = errors.New("failed decoding input file data")
)

func ExtractXmlData(inFilePath string) (currency.CurrencyList, error) {
        inFile, err := openInFile(inFilePath)
        if err != nil {
                return nil, errors.Join(errOpenInpFileFailed, err)
        }
        defer inFile.Close()

        decoder := createXmlDecoder(inFile)

        data, err := decodeXmlFile(decoder)
        if err != nil {
                return nil, errors.Join(errDecodeInpFileFailed, err)
        }

        return data, nil
}

func openInFile(inFilePath string) (*os.File, error) {
        inFile, err := os.Open(inFilePath)
        if err != nil {
                var pathErr *fs.PathError
                if errors.As(err, &pathErr) {
                        err = fmt.Errorf("failed to open config file path: %s",
                                         pathErr.Path)
                }

                return nil, err
        }

        return inFile, nil
}

// separate function for flexibility
func createXmlDecoder(inFile io.Reader) *xml.Decoder {
        decoder := xml.NewDecoder(inFile)
        decoder.CharsetReader = charset.NewReaderLabel
        return decoder
}

var (
        errRecordsDecodeFailed = errors.New("failed decoding an xml currency records")
        errValidFailed         = errors.New("failed validating decoded xml records")

)

func decodeXmlFile(decoder *xml.Decoder) (currency.CurrencyList, error) {
        if decoder == nil {
                // `panic` is used here as an assertion: it can be
                // triggered only by a critical memory fault or
                // because of a developer's mistake

                panic("failed decoding xml file data")
        }

        scheme := currency.NewScheme()
        err := decoder.Decode(&scheme)
        if err != nil {
                return nil, errors.Join(errRecordsDecodeFailed, err)
        }

        err = validator.New().Struct(scheme)
        if err != nil {
                return nil, errors.Join(errValidFailed, err)
        }

        return scheme.List, nil
}
