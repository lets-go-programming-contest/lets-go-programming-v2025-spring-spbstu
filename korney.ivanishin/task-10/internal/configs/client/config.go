package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type ConfigParams struct {
        ServicePort string `yaml:"service_port" validate:"required"`
}

func new() ConfigParams {
        return ConfigParams{ ServicePort: `` }
}

var (
        errConfReadFailed    = errors.New("failed reading config file data")
        errConfProcFailed    = errors.New("failed processing config file data")
        errUnmrashalFailed   = errors.New("failed unmarshalling")
        errDecodeValidFailed = errors.New("decoded data validation failed")
)

func GetConfigParams() (ConfigParams, error) {
        confFilePath := `config/client/config.yaml`

        confFileContents, err := readInFile(confFilePath)
        if err != nil {
                zeroConfigParams := new()
                return zeroConfigParams, errors.Join(errConfReadFailed, err)
        }

        configParams, err := decodeConfFileData(confFileContents)
        if err != nil {
                zeroConfigParams := new()
                return zeroConfigParams, errors.Join(errConfProcFailed, err)
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
                zeroConfigParams := new()
                return zeroConfigParams, errors.Join(errUnmrashalFailed, err)
        }

        err = validator.New().Struct(configParams)
        if err != nil {
                zeroConfigParams := new()
                return zeroConfigParams, errors.Join(errDecodeValidFailed, err)
        }

        return configParams, nil
}
