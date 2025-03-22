package main

import (
	"errors"

	"github.com/quaiion/go-practice/convertation/internal/config"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/encase"
	"github.com/quaiion/go-practice/convertation/internal/dataProcessing/extract"
)

var (
        errConfigFailed =  errors.New("configuration failed")
        errExtractFailed = errors.New("data extraction failed")
        errEncaseFailed =  errors.New("data encasement failed")
)

func main() {
        inFilePath, outFilePath, err := config.GetIOFilePaths()
        if err != nil {
                panic(errors.Join(errConfigFailed, err))
        }

        data, err := extract.ExtractXmlData(inFilePath)
        if err != nil {
                panic(errors.Join(errExtractFailed, err))
        }

        err = encase.EncaseJsonData(outFilePath, data)
        if err != nil {
                panic(errors.Join(errEncaseFailed, err))
        }
}
