package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dmitriy.rumyantsev/task-3/internal/models"
	my_validator "github.com/dmitriy.rumyantsev/task-3/internal/validator"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/net/html/charset"
)

// Reads and parses currency data in XML format
func ReadValCurs(path string) (*models.ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML file: %w", err)
	}

	// Create decoder with charset support
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs models.ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML data: %w", err)
	}

	return &valCurs, nil
}

// Parses a string with a number where the decimal separator can be a comma
func ParseLocaleFloat(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", ".")

	return strconv.ParseFloat(s, 64)
}

// Processes currency data, converts values and sorts by descending order
func ProcessValutesData(valCurs *models.ValCurs) ([]models.OutputValute, error) {
	var outputValutes []models.OutputValute
	validate := validator.New()

	for _, valute := range valCurs.Valutes {
		value, err := ParseLocaleFloat(valute.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse value '%s': %w", valute.Value, err)
		}

		outputValute := models.OutputValute{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		}

		err = validate.Struct(outputValute)
		if err != nil {
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				return nil, fmt.Errorf("validation error for currency %s: %v",
					valute.CharCode, my_validator.FormatValidationErrors(validationErrors))
			}
			return nil, fmt.Errorf("unknown validation error: %w", err)
		}

		outputValutes = append(outputValutes, outputValute)
	}

	// Sort by value in descending order
	sort.Slice(outputValutes, func(i, j int) bool {
		return outputValutes[i].Value > outputValutes[j].Value
	})

	return outputValutes, nil
}
