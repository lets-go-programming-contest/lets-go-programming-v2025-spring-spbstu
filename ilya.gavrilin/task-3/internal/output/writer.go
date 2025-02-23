package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// WriteToFile saves processed data to a JSON file.
func WriteToFile(path string, data interface{}) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", path, err)
	}
	defer file.Close()

	// Encode to JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to write JSON: %w", err)
	}

	return nil
}