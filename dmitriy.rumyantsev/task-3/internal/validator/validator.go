package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Formats validation errors into a readable form
func FormatValidationErrors(errors validator.ValidationErrors) string {
	var errorMessages []string

	for _, err := range errors {
		switch err.Tag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("field '%s' is required", err.Field()))
		case "len":
			errorMessages = append(errorMessages, fmt.Sprintf("field '%s' must be exactly %s characters long", err.Field(), err.Param()))
		case "gt":
			errorMessages = append(errorMessages, fmt.Sprintf("field '%s' must be greater than %s", err.Field(), err.Param()))
		case "file":
			errorMessages = append(errorMessages, fmt.Sprintf("file '%s' does not exist", err.Value()))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("field '%s' validation failed: %s", err.Field(), err.Tag()))
		}
	}

	return strings.Join(errorMessages, "; ")
}
