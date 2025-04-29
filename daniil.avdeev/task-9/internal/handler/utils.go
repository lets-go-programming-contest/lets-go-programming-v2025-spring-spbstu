package handler

import (
  "errors"
	"encoding/json"
	"net/http"

	"github.com/realFrogboy/task-9/internal/contacts"
)

func checkContact(contact contacts.Contact) error {
	if contact.Name == "" {
		return errors.New("Name is required")
	}

	if contact.Phone == "" {
		return errors.New("Phone is required")
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
