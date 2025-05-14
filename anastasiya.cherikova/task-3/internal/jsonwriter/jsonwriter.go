package jsonwriter

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"task-3/internal/currency"
)

var (
	ErrCreateDir  = errors.New("directory create error")
	ErrCreateFile = errors.New("file create error")
	ErrJSONEncode = errors.New("json encode error")
)

// WriteJSON saves sorted currencies to a JSON file
func WriteJSON(currencies []currency.Currency, outputPath string) error {
	// Create a directory if it does not exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateDir, err)
	}

	// Creating a file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateFile, err)
	}
	defer file.Close()

	// Formatted JSON
	encoder := json.NewEncoder(file)
	// SetIndent configures JSON formatting for better human readability.
	// First argument "" means no line prefix.
	// Second argument "    " uses 4 spaces for indentation (common standard).
	// This produces nicely formatted JSON instead of a single-line output.
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("%w: %v", ErrJSONEncode, err)
	}
	return nil
}
