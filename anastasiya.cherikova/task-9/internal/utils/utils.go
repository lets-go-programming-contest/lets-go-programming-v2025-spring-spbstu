package utils

import (
	"regexp"
	"strings"
)

// Validate phone number format
func IsValidPhone(phone string) bool {
	// Allow formats: +79991112233, 89991112233, 9991112233
	normalized := strings.TrimSpace(phone)
	pattern := `^\+?\d{10,15}$`
	matched, _ := regexp.MatchString(pattern, normalized)
	return matched
}
