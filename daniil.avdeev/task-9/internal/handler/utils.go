package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/realFrogboy/task-9/internal/contacts"
)

func checkContact(contact contacts.Contact) error {
	nameRegex, err := regexp.Compile(`[A-Za-z]+(\s[A-Za-z]+)*`)
	if err != nil {
		return errors.New("Faulty name regex")
	}

	if !nameRegex.MatchString(contact.Name) {
		return errors.New("Invalid name format")
	}

	phoneRegex, err := regexp.Compile(`(^8|7|\+7)((\d{10})|(\s\(\d{3}\)\s\d{3}\s\d{2}\s\d{2}))`)
	if err != nil {
		return errors.New("Faulty phone regex")
	}

	if !phoneRegex.MatchString(contact.Phone) {
		return errors.New("Invalid phone number format")
	}
	return nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, contacts.ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}
