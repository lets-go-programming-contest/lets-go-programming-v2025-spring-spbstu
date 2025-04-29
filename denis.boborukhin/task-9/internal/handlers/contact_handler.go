package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"unicode/utf8"

	"github.com/denisboborukhin/contact_manager/internal/contact/manager"
	"github.com/gorilla/mux"
)

type ContactHandler struct {
	manager *manager.Manager
}

func NewContactHandler(manager *manager.Manager) *ContactHandler {
	return &ContactHandler{
		manager: manager,
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

type ContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string, details string) {
	writeJSON(w, status, ErrorResponse{
		Error:   message,
		Details: details,
	})
}

// validateContact performs basic validation of contact data
func validateContact(name, phone string) error {
	// Check name
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if utf8.RuneCountInString(name) < 2 {
		return fmt.Errorf("name must be at least 2 characters long")
	}
	if utf8.RuneCountInString(name) > 50 {
		return fmt.Errorf("name must not exceed 50 characters")
	}

	// Check phone
	if phone == "" {
		return fmt.Errorf("phone number is required")
	}
	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{10,14}$`)
	if !phoneRegex.MatchString(phone) {
		return fmt.Errorf("phone number must be in format +7XXXXXXXXXX")
	}

	return nil
}

// GetContacts handles GET /contacts request
func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.manager.ListContacts(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to retrieve contacts", err.Error())
		return
	}
	writeJSON(w, http.StatusOK, contacts)
}

// GetContact handles GET /contacts/{id} request
func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	contact, err := h.manager.GetContact(r.Context(), id)
	if err != nil {
		if err == manager.ErrContactNotFound {
			writeError(w, http.StatusNotFound, "Contact not found", "")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to get contact", err.Error())
		return
	}
	writeJSON(w, http.StatusOK, contact)
}

// CreateContact handles POST /contacts request
func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
	defer r.Body.Close()

	if err := validateContact(req.Name, req.Phone); err != nil {
		writeError(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	contact, err := h.manager.CreateContact(r.Context(), req.Name, req.Phone)
	if err != nil {
		if err == manager.ErrInvalidInput {
			writeError(w, http.StatusBadRequest, "Invalid input", err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to create contact", err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, contact)
}

// UpdateContact handles PUT /contacts/{id} request
func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
	defer r.Body.Close()

	if err := validateContact(req.Name, req.Phone); err != nil {
		writeError(w, http.StatusBadRequest, "Validation failed", err.Error())
		return
	}

	contact, err := h.manager.UpdateContact(r.Context(), id, req.Name, req.Phone)
	if err != nil {
		switch err {
		case manager.ErrContactNotFound:
			writeError(w, http.StatusNotFound, "Contact not found", "")
		case manager.ErrInvalidInput:
			writeError(w, http.StatusBadRequest, "Invalid input", err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "Failed to update contact", err.Error())
		}
		return
	}

	writeJSON(w, http.StatusOK, contact)
}

// DeleteContact handles DELETE /contacts/{id} request
func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.manager.DeleteContact(r.Context(), id); err != nil {
		if err == manager.ErrContactNotFound {
			writeError(w, http.StatusNotFound, "Contact not found", "")
			return
		}
		writeError(w, http.StatusInternalServerError, "Failed to delete contact", err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
