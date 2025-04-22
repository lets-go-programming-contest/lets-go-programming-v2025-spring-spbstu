package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"task-9/internal/model"
	"task-9/internal/storage"
)

type Handler struct {
	store *storage.ContactStore
}

const (
	ErrInvalidRequest = 400
	ErrNotFound       = 404
	ErrDuplicateEntry = 409
	ErrServerError    = 500
)

// Phone number validation
// +XXXXXXXXXXX or 8XXXXXXXXXX
var phoneRegex = regexp.MustCompile(`^(\+\d{1,3}|8)\d{10}$`)

func NewHandler(store *storage.ContactStore) *Handler {
	return &Handler{store: store}
}

// Request type		endpoint
//
//	  GET          /contacts
//	 POST          /contacts
//	OPTIONS        /contacts <- optional one
func (h *Handler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.getAllContacts(w, r)
	case http.MethodPost:
		h.createContact(w, r)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		w.WriteHeader(http.StatusOK)
	default:
		h.errorResponse(w, http.StatusMethodNotAllowed, ErrInvalidRequest, "Method not allowed", "Use GET or POST")
	}
}

// Request type		endpoint
//
//	  GET          /contacts/{id}
//	  PUT          /contacts/{id}
//	DELETE         /contacts/{id}
//	OPTIONS        /contacts/{id} <- optional one
func (h *Handler) HandleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := strings.TrimPrefix(r.URL.Path, "/contacts/")
	if id == "" {
		h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Invalid request", "Contact ID required")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getContactByID(w, id)
	case http.MethodPut:
		h.updateContact(w, r, id)
	case http.MethodDelete:
		h.deleteContact(w, id)
	case http.MethodOptions:
		w.Header().Set("Allow", "GET, PUT, DELETE, OPTIONS")
		w.WriteHeader(http.StatusOK)
	default:
		h.errorResponse(w, http.StatusMethodNotAllowed, ErrInvalidRequest, "Method not allowed", "Use GET, PUT or DELETE")
	}
}

func (h *Handler) getAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.store.GetAll()
	if err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error retrieving contacts", err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error encoding response", err.Error())
	}
}

func (h *Handler) getContactByID(w http.ResponseWriter, id string) {
	contact, err := h.store.GetByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			h.errorResponse(w, http.StatusNotFound, ErrNotFound, "Contact not found", "No contact with the given ID exists")
		} else {
			h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error retrieving contact", err.Error())
		}
		return
	}

	if err := json.NewEncoder(w).Encode(contact); err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error encoding response", err.Error())
	}
}

func (h *Handler) createContact(w http.ResponseWriter, r *http.Request) {
	var contact model.Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Invalid request body", "Unable to parse JSON")
		return
	}

	if err := validateContact(contact); err != nil {
		h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Validation error", err.Error())
		return
	}

	createdContact, err := h.store.Create(contact)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrDuplicateContact):
			h.errorResponse(w, http.StatusConflict, ErrDuplicateEntry, "Duplicate contact", "A contact with this name and phone already exists")
		case errors.Is(err, storage.ErrInvalidData):
			h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Invalid data", "Name and phone are required")
		default:
			h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error creating contact", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdContact); err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error encoding response", err.Error())
	}
}

func (h *Handler) updateContact(w http.ResponseWriter, r *http.Request, id string) {
	var contact model.Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Invalid request body", "Unable to parse JSON")
		return
	}

	if err := validateContact(contact); err != nil {
		h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Validation error", err.Error())
		return
	}

	updatedContact, err := h.store.Update(id, contact)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			h.errorResponse(w, http.StatusNotFound, ErrNotFound, "Contact not found", "No contact with the given ID exists")
		case errors.Is(err, storage.ErrDuplicateContact):
			h.errorResponse(w, http.StatusConflict, ErrDuplicateEntry, "Duplicate contact", "A contact with this name and phone already exists")
		case errors.Is(err, storage.ErrInvalidData):
			h.errorResponse(w, http.StatusBadRequest, ErrInvalidRequest, "Invalid data", "Name and phone are required")
		default:
			h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error updating contact", err.Error())
		}
		return
	}

	if err := json.NewEncoder(w).Encode(updatedContact); err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error encoding response", err.Error())
	}
}

func (h *Handler) deleteContact(w http.ResponseWriter, id string) {
	contact, err := h.store.GetByID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			h.errorResponse(w, http.StatusNotFound, ErrNotFound, "Contact not found", "No contact with the given ID exists")
		} else {
			h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error retrieving contact", err.Error())
		}
		return
	}

	if err := h.store.Delete(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			h.errorResponse(w, http.StatusNotFound, ErrNotFound, "Contact not found", "No contact with the given ID exists")
		} else {
			h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error deleting contact", err.Error())
		}
		return
	}

	response := struct {
		Message string        `json:"message"`
		Contact model.Contact `json:"contact"`
	}{
		Message: "Contact successfully deleted",
		Contact: contact,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.errorResponse(w, http.StatusInternalServerError, ErrServerError, "Error encoding response", err.Error())
	}
}

func (h *Handler) errorResponse(w http.ResponseWriter, httpStatus, code int, message, details string) {
	w.WriteHeader(httpStatus)

	resp := model.ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}

	json.NewEncoder(w).Encode(resp)
}

func validateContact(c model.Contact) error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Phone == "" {
		return errors.New("phone number is required")
	}

	if !phoneRegex.MatchString(c.Phone) {
		return errors.New("invalid phone format (use +XXXXXXXXXXX or 8XXXXXXXXXX)")
	}

	return nil
}
