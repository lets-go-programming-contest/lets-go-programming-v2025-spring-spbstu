package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"task-9/internal/models"
	"task-9/internal/utils"

	"github.com/gorilla/mux"
)

// Return all contacts
func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contacts, err := models.GetAllContacts(db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve contacts")
			return
		}
		respondWithJSON(w, http.StatusOK, contacts)
	}
}

// Return a single contact by ID
func GetContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		contact, err := models.GetContactByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				respondWithError(w, http.StatusNotFound, "Contact not found")
			} else {
				respondWithError(w, http.StatusInternalServerError, "Failed to retrieve contact")
			}
			return
		}
		respondWithJSON(w, http.StatusOK, contact)
	}
}

// Create a new contact
func CreateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.Contact
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if !utils.IsValidPhone(c.Phone) {
			respondWithError(w, http.StatusBadRequest, "Invalid phone number format")
			return
		}

		id, err := models.CreateContact(db, c)
		if err != nil {
			if err.Error() == "UNIQUE constraint failed: contacts.phone" {
				respondWithError(w, http.StatusConflict, "Phone number already exists")
			} else {
				respondWithError(w, http.StatusInternalServerError, "Failed to create contact")
			}
			return
		}

		c.ID = int(id)
		respondWithJSON(w, http.StatusCreated, c)
	}
}

// Update an existing contact
func UpdateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		var c models.Contact
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if !utils.IsValidPhone(c.Phone) {
			respondWithError(w, http.StatusBadRequest, "Invalid phone number format")
			return
		}

		if err := models.UpdateContact(db, id, c); err != nil {
			if err.Error() == "UNIQUE constraint failed: contacts.phone" {
				respondWithError(w, http.StatusConflict, "Phone number already exists")
			} else {
				respondWithError(w, http.StatusInternalServerError, "Failed to update contact")
			}
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Contact updated successfully"})
	}
}

// Delete a contact
func DeleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		if err := models.DeleteContact(db, id); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete contact")
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Contact deleted successfully"})
	}
}

// Helper functions
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
