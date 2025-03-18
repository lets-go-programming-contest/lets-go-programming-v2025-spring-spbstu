package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"task-3/internal/currency"

	"golang.org/x/text/encoding/charmap"
)

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"` // Значение как строка
}

type ValCurs struct {
	Currencies []xmlCurrency `xml:"Valute"`
}

func ParseXML(filePath string) []currency.Currency {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Ошибка открытия файла: %v", err))
	}
	defer file.Close()

	// Создаем XML-декодер с обработкой кодировки
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			// Преобразуем Windows-1251, UTF-8
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("неподдерживаемая кодировка: %s", charset)
	}

	// Парсим XML
	var valCurs ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(fmt.Sprintf("Ошибка декодирования XML: %v", err))
	}

	// Конвертируем значения в float64
	result := make([]currency.Currency, 0, len(valCurs.Currencies))
	for _, c := range valCurs.Currencies {
		valueStr := strings.Replace(c.ValueStr, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(fmt.Sprintf("Ошибка конвертации '%s': %v", c.ValueStr, err))
		}

		result = append(result, currency.Currency{
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Value:    value,
		})
	}

	return result
}
