package encase

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/currency"
)

func EncaseJsonData(outFilePath string, currList currency.CurrencyList) error {
        if currList == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed encasing/writing data")
        }

        err := prepareOutputEnv(outFilePath)
        if err != nil {
                return fmt.Errorf("failed preparing the output environment // %w",
                                  err)
        }

        currList = transformData(currList)

        data, err := json.MarshalIndent(currList, ``, ` `)
        if err != nil {
                return fmt.Errorf("failed encoding (marshalling) output data // %w",
                                  err)
        }

        err = writeJsonData(data, outFilePath)
        if err != nil {
                return fmt.Errorf("failed writing output json data // %w",
                                  err)
        }

        return nil
}

/** separate function just for flexibility */
func prepareOutputEnv(filePath string) error {
        err := os.MkdirAll(filepath.Dir(filePath), 0644)
        if err != nil {
                return fmt.Errorf("failed to create specified directories // %w",
                                  err)
        }

        return nil
}

/** separate function just for flexibility */
func transformData(currList currency.CurrencyList) currency.CurrencyList {
        if currList == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed transforming decoded data")
        }

        currList.Sort()
        return currList
}

/** separate function just for flexibility */
func writeJsonData(data []byte, outFilePath string) error {
        if data == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed writing encased data")
        }

        err := os.WriteFile(outFilePath, data, 0644)
        if err != nil {
                return fmt.Errorf("failed writing output data to the file // %w",
                                  err)
        }

        return nil
}
