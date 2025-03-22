package Config

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type ConfigParams struct {
        NRequesters uint32 `yaml:"n_requesters" validate:"required"`
        ReqRange    uint32 `yaml:"request_range" validate:"required"`
        CacheCap    uint32 `yaml:"cache_capacity" validate:"required"`
        NRequests   uint32 `yaml:"n_requests" validate:"required"`
}

func GetConfigParams() (ConfigParams, error) {
        confFilePath := parseConfFilePathFlag()

        confFileContents, err := readInFile(confFilePath)
        if err != nil {
                zeroConfigParams := ConfigParams{0, 0, 0, 0}
                return zeroConfigParams, fmt.Errorf("failed reading config file data: %w",
                                                    err)
        }

        configParams, err := decodeConfFileData(confFileContents)
        if err != nil {
                zeroConfigParams := ConfigParams{0, 0, 0, 0}
                return zeroConfigParams, fmt.Errorf("failed processing config file data: %w",
                                                    err)
        }

        return configParams, nil
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

func decodeConfFileData(confFileContents []byte) (ConfigParams, error) {
        if confFileContents == nil {
                //  `panic` is used here as an assertion: it can be
                //  triggered only by a critical memory fault or
                //  because of a developer's mistake

                panic("failed while opening a file / storing its contents")
        }

        var configParams ConfigParams

        err := yaml.Unmarshal(confFileContents, &configParams)
        if err != nil {
                zeroConfigParams := ConfigParams{0, 0, 0, 0}
                return zeroConfigParams, fmt.Errorf("failed unmarshalling: %w", err)
        }

        err = validator.New().Struct(configParams)
        if err != nil {
                zeroConfigParams := ConfigParams{0, 0, 0, 0}
                return zeroConfigParams, fmt.Errorf("decoded data validation failed: %w",
                                                    err)
        }

        return configParams, nil
}
