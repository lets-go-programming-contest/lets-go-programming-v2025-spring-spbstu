package config

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var (
        errConfReadFailed = errors.New("failed reading config file data")
        errConfProcFailed = errors.New("failed processing config file data")       
)

func GetIOFilePaths() (string, string, error) {
        confFilePath := parseConfFilePathFlag()

        confFileContents, err := readInFile(confFilePath)
        if err != nil {
                return ``, ``, errors.Join(errConfReadFailed, err)
        }

        inFilePath, outFilePath, err := decodeConfFileData(confFileContents)
        if err != nil {
                return ``, ``, errors.Join(errConfProcFailed, err)
        }

        return inFilePath, outFilePath, nil
}

func readInFile(filePath string) ([]byte, error) {
        fileData, err := os.ReadFile(filePath)
        if err != nil {
                var pathErr *fs.PathError
                if errors.As(err, &pathErr) {
                        err = fmt.Errorf("failed to open config file path: %s",
                                         pathErr.Path)
                }

                return nil, err
        }

        return fileData, nil
}

func parseConfFilePathFlag() (string) {
        var pathStr string
        flag.StringVar(&pathStr, "config", "config/config.yml", "config file path")
        flag.Parse()

        return pathStr
}

var (
        errUnmrashalFailed =   errors.New("failed unmarshalling")
        errDecodeValidFailed = errors.New("decoded data validation failed")
)

type fileNamesParsed struct {
        InFile  string `yaml:"input-file" validate:"required"`
        OutFile string `yaml:"output-file" validate:"required"`
}

func decodeConfFileData(confFileContents []byte) (string, string, error) {
        var parsed fileNamesParsed

        err := yaml.Unmarshal(confFileContents, &parsed)
        if err != nil {
                return ``, ``, errors.Join(errUnmrashalFailed, err)
        }

        err = validator.New().Struct(parsed)
        if err != nil {
                return ``, ``, errors.Join(errDecodeValidFailed, err)
        }

        return parsed.InFile, parsed.OutFile, nil
}
