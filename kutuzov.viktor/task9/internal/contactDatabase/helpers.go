package contactdatabase

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func isValidContact(c Contact) bool {
	if c.Name == "" || c.Phone == "" {
		return false
	}
	return isValidPhone(c.Phone)
}

func isValidPhone(phone string) bool {
	matched, _ := regexp.MatchString(`^\+?[0-9]{10,15}$`, phone)
	return matched
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
