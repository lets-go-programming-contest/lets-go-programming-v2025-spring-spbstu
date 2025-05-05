package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/realFrogboy/task-9/internal/contacts"
	"github.com/realFrogboy/task-9/internal/db"

	"github.com/gorilla/mux"
)

type ContactHandler struct {
  db *db.SQLiteStorage
}

func NewContactHandler(db *db.SQLiteStorage) *ContactHandler {
	return &ContactHandler{db: db}
}

func (h *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.db.GetAllContacts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, contacts)
}

func (h *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	contact, err := h.db.GetContactByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if contact == nil {
		respondWithError(w, http.StatusNotFound, "Contact not found")
		return
	}

	respondWithJSON(w, http.StatusOK, contact)
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact contacts.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

  if err := checkContact(contact); err != nil {
    respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.db.CreateContact(contact)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	dbContact := contacts.DBContact{
		ID:    id,
		Name:  contact.Name,
		Phone: contact.Phone,
	}

	respondWithJSON(w, http.StatusCreated, dbContact)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	var contact contacts.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

  if err := checkContact(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

  if err := h.db.UpdateContact(id, contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	dbContact := contacts.DBContact{
		ID:    id,
		Name:  contact.Name,
		Phone: contact.Phone,
	}

	respondWithJSON(w, http.StatusOK, dbContact)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid contact ID")
		return
	}

	if err := h.db.DeleteContact(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
