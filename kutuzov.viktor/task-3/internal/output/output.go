package output

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vktr-ktzv/task3/internal/dataHandler"
)

func SaveJSON(outputFile string, currencies []dataHandler.Currency) {
	outputDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Error creating dir: %v", err))
	}

	file, err := os.Create(outputFile)
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(currencies); err != nil {
		panic(fmt.Sprintf("Error encoding JSON: %v", err))
	}
}
