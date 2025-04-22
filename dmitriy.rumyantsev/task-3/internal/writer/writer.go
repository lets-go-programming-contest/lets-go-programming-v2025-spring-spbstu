package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dmitriy.rumyantsev/task-3/internal/models"
)

// Saves processed data as JSON
func SaveAsJSON(outputValutes []models.OutputValute, path string) error {
	// Check if directory exists and create it if necessary
	dir := filepath.Dir(path)
	if dir != "." {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory for output file: %w", err)
		}
	}

	// Convert data to JSON
	jsonData, err := json.MarshalIndent(outputValutes, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	// Write JSON to file
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}
