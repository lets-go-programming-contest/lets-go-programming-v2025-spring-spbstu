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

func GetConfigParams() (uint32, uint32, uint32, uint32, error) {
        confFilePath := parseConfFilePathFlag()

        confFileContents, err := readInFile(confFilePath)
        if err != nil {
                return 0, 0, 0, 0, fmt.Errorf("failed reading config file data: %w",
                                              err)
        }

        nRequesters, reqRange, cacheCap, nRequests, err := decodeConfFileData(confFileContents)
        if err != nil {
                return 0, 0, 0, 0, fmt.Errorf("failed processing config file data: %w",
                                              err)
        }

        return nRequesters, reqRange, cacheCap, nRequests, nil
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

type configParamsParsed struct {
        NRequesters uint32 `yaml:"n_requesters" validate:"required"`
        ReqRange    uint32 `yaml:"request_range" validate:"required"`
        CacheCap    uint32 `yaml:"cache_capacity" validate:"required"`
        NRequests   uint32 `yaml:"n_requests" validate:"required"`
}

func decodeConfFileData(confFileContents []byte) (uint32, uint32, uint32, uint32, error) {
        if confFileContents == nil {
                /** 
                 * `panic` is used here as an assertion: it can be
                 * triggered only by a critical memory fault or
                 * because of a developer's mistake
                 */
                panic("failed while opening a file / storing its contents")
        }

        var parsed configParamsParsed

        err := yaml.Unmarshal(confFileContents, &parsed)
        if err != nil {
                return 0, 0, 0, 0, fmt.Errorf("failed unmarshalling: %w", err)
        }

        err = validator.New().Struct(parsed)
        if err != nil {
                return 0, 0, 0, 0, fmt.Errorf("decoded data validation failed: %w",
                                           err)
        }

        return parsed.NRequesters, parsed.ReqRange, parsed.CacheCap, parsed.NRequests, nil
}
