package main

import (
	"fmt"

	"github.com/quaiion/go-practice/convertation/internal/config"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/encase"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/extract"
)

func main() {
        inFilePath, outFilePath, err := config.GetIOFilePaths()
        if err != nil {
                err = fmt.Errorf("configuration failed // %w", err)
                fmt.Println(err)
                return
        }

        data, err := extract.ExtractXmlData(inFilePath)
        if err != nil {
                err = fmt.Errorf("data extraction failed // %w", err)
                fmt.Println(err)
                return
        }

        err = encase.EncaseJsonData(outFilePath, data)
        if err != nil {
                err = fmt.Errorf("data encasement failed // %w", err)
                fmt.Println(err)
                return
        }
}
