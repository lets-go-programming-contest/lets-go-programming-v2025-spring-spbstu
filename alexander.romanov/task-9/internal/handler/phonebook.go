package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alexander-romanov-edu/phonebook/internal/service"
	"github.com/alexander-romanov-edu/phonebook/pkg/models"
)

type PhonebookHandler struct {
	service service.PhonebookService
}

func NewPhonebookHandler(s service.PhonebookService) *PhonebookHandler {
	return &PhonebookHandler{service: s}
}

func (h *PhonebookHandler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		contacts, err := h.service.GetAllContacts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	case http.MethodPost:
		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if contact.Name == "" || contact.Phone == "" {
			http.Error(w, "name and phone are required", http.StatusBadRequest)
			return
		}

		if err := h.service.AddContact(contact); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(contact)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PhonebookHandler) HandleSingleContact(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		contact, err := h.service.GetContact(id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Contact not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)
	case http.MethodPut:
		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		contact.ID = id
		if err := h.service.UpdateContact(contact); err != nil {
			if err.Error() == "contact not found" {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(contact)
	case http.MethodDelete:
		if err := h.service.DeleteContact(id); err != nil {
			if err.Error() == "contact not found" {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Contact with ID %s deleted successfully", id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
