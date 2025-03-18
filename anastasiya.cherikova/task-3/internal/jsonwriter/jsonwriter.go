package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"task-3/internal/currency"
)

// Запись отсортированных данных в JSON
func WriteJSON(currencies []currency.Currency, outputPath string) {
	// Создание директории, если её нет
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Ошибка создания директории: %v", err))
	}

	// Создание файла
	file, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка создания файла: %v", err))
	}
	defer file.Close()

	// Форматированный JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(currencies); err != nil {
		panic(fmt.Sprintf("Ошибка записи JSON: %v", err))
	}
}
