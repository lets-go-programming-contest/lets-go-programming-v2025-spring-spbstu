package extract

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/currency"
	"golang.org/x/net/html/charset"
)

func ExtractXmlData(inFilePath string) (currency.CurrencyList, error) {
        inFile, err := openInFile(inFilePath)
        if err != nil {
                return nil, fmt.Errorf("failed opening input file // %w",
                                       err)
        }
        defer inFile.Close()

        decoder := createXmlDecoder(inFile)

        data, err := decodeXmlFile(decoder)
        if err != nil {
                return nil, fmt.Errorf("failed decoding input file data // %w",
                                       err)
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

/** separate function for flexibility */
func createXmlDecoder(inFile io.Reader) *xml.Decoder {
        decoder := xml.NewDecoder(inFile)
        decoder.CharsetReader = charset.NewReaderLabel
        return decoder
}

func decodeXmlFile(decoder *xml.Decoder) (currency.CurrencyList, error) {
        if decoder == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed decoding xml file data")
        }

        var currList []currency.Currency

        for token, err := decoder.Token() ; token != nil ; token, err = decoder.Token() {
                if err != nil {
                        return nil, fmt.Errorf("failed parsing a token from input file data // %w",
                                               err)
                }

                switch tokenType := token.(type) {
                case xml.StartElement:
                        if tokenType.Name.Local != `Valute` {
                                continue
                        }

                        var curr currency.Currency
                        err = decoder.DecodeElement(&curr, &tokenType)
                        if err != nil {
                                return nil, fmt.Errorf("failed decoding an xml currency record // %w",
                                                       err)
                        }

                        err = validator.New().Struct(curr)
                        if err != nil {
                                /**
                                 * we only validate the "required" condition and
                                 * just discard elements that dont satisfy it,
                                 * so there is no need to pass this error upwards
                                 */
                                 continue
                        }

                        curr, err = translateValueStrToValue(curr)
                        if err != nil {
                                return nil, fmt.Errorf("failed translating currency '%s' value string to value // %w",
                                                       curr.CharCode, err)
                        }

                        currList = append(currList, curr)
                }
        }

        return currList, nil
}

func translateValueStrToValue(curr currency.Currency) (currency.Currency, error) {
        curr.Value = strings.ReplaceAll(curr.Value, `,`, `.`)

        var err error = nil
        curr.FPValue, err = strconv.ParseFloat(curr.Value, 64)
        if err != nil {
                return curr, fmt.Errorf("failed converting a 'Value' record to float // %w",
                                        err)
        }

        return curr, nil
}
