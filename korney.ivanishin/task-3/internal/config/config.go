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

func GetIOFilePaths() (string, string, error) {
        confFilePath, isDefaultPath := parseConfFilePathFlag()

        confFileContents, err := readInFile(confFilePath)
        if err != nil {
                if isDefaultPath {
                        return ``, ``, fmt.Errorf("failed reading (default) config file data // %w",
                                                  err)
                } else {
                        return ``, ``, fmt.Errorf("failed reading config file data // %w",
                                                  err)
                }
        }

        inFilePath, outFilePath, err := decodeConfFileData(confFileContents)
        if err != nil {
                return ``, ``, fmt.Errorf("failed processing config file data // %w",
                                          err)
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

func parseConfFilePathFlag() (string, bool) {
        var pathStr string
        flag.StringVar(&pathStr, "config", "config/config.yml", "config file path")
        flag.Parse()
        return pathStr, flag.Parsed()
}

type fileNamesParsed struct {
        InFile  string `yaml:"input-file" validate:"required"`
        OutFile string `yaml:"output-file" validate:"required"`
}

func decodeConfFileData(confFileContents []byte) (string, string, error) {
        if confFileContents == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed while opening a file / storing its contents")
        }

        var parsed fileNamesParsed

        err := yaml.Unmarshal(confFileContents, &parsed)
        if err != nil {
                return ``, ``, fmt.Errorf("failed unmarshalling // %w", err)
        }

        err = validator.New().Struct(parsed)
        if err != nil {
                return ``, ``, fmt.Errorf("decoded data validation failed // %w",
                                          err)
        }

        return parsed.InFile, parsed.OutFile, nil
}
