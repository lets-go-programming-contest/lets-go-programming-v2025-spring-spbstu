package contact

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllContacts(w, r)
	case http.MethodPost:
		h.createContact(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleSingleContact(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	if id, err := strconv.Atoi(parts[2]); err == nil {
		switch r.Method {
		case http.MethodGet:
			h.getContactByID(w, r, id)
		case http.MethodPut:
			h.updateContact(w, r, id)
		case http.MethodDelete:
			h.deleteContactByID(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		name := parts[2]
		switch r.Method {
		case http.MethodGet:
			h.getContactByName(w, r, name)
		case http.MethodDelete:
			h.deleteContactByName(w, r, name)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *Handler) getContactByName(w http.ResponseWriter, r *http.Request, name string) {
	contact, err := h.service.GetContactByName(name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve contact", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact retrieved successfully",
		Data: ContactResponse{
			ID:    contact.ID,
			Name:  contact.Name,
			Phone: contact.Phone,
		},
	})
}

func (h *Handler) deleteContactByName(w http.ResponseWriter, r *http.Request, name string) {
	if err := h.service.DeleteContactByName(name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact deleted successfully",
	})
}

func (h *Handler) deleteContactByID(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.DeleteContactByID(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact deleted successfully",
	})
}
func (h *Handler) getAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		http.Error(w, "Failed to retrieve contacts", http.StatusInternalServerError)
		return
	}

	var response []ContactResponse
	for _, c := range contacts {
		response = append(response, ContactResponse{
			ID:    c.ID,
			Name:  c.Name,
			Phone: c.Phone,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contacts retrieved successfully",
		Data:    response,
	})
}

func (h *Handler) getContactByID(w http.ResponseWriter, r *http.Request, id int) {
	contact, err := h.service.GetContactByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve contact", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact retrieved successfully",
		Data: ContactResponse{
			ID:    contact.ID,
			Name:  contact.Name,
			Phone: contact.Phone,
		},
	})
}

func (h *Handler) createContact(w http.ResponseWriter, r *http.Request) {
	var c Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdContact, err := h.service.CreateContact(c)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			http.Error(w, "Phone number already exists", http.StatusConflict)
		} else {
			http.Error(w, fmt.Sprintf("Failed to create contact: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact created successfully",
		Data: ContactResponse{
			ID:    createdContact.ID,
			Name:  createdContact.Name,
			Phone: createdContact.Phone,
		},
	})
}

func (h *Handler) updateContact(w http.ResponseWriter, r *http.Request, id int) {
	var c Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedContact, err := h.service.UpdateContact(id, c)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else if strings.Contains(err.Error(), "duplicate key") {
			http.Error(w, "Phone number already exists", http.StatusConflict)
		} else {
			http.Error(w, fmt.Sprintf("Failed to update contact: %v", err), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact updated successfully",
		Data: ContactResponse{
			ID:    updatedContact.ID,
			Name:  updatedContact.Name,
			Phone: updatedContact.Phone,
		},
	})
}

func (h *Handler) deleteContact(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.DeleteContactByID(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Contact not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "Contact deleted successfully",
	})
}