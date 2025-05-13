package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/realFrogboy/task-9/internal/contacts"
	"github.com/realFrogboy/task-9/internal/db"

	"github.com/gorilla/mux"
)

type ContactHandler struct {
	db     *db.SQLiteStorage
	router *mux.Router
}

func NewContactHandler(dbPath string) (*ContactHandler, error) {
	dbStorage, err := db.NewSQLiteStorage(dbPath)
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	return &ContactHandler{db: dbStorage, router: router}, nil
}

func (h *ContactHandler) Delete() {
	h.db.Close()
}

func (h *ContactHandler) Run(serverPath string) {
	h.router.HandleFunc("/contacts", h.GetAllContacts).Methods(http.MethodGet)
	h.router.HandleFunc("/contacts/{id}", h.GetContact).Methods(http.MethodGet)
	h.router.HandleFunc("/contacts", h.CreateContact).Methods(http.MethodPost)
	h.router.HandleFunc("/contacts/{id}", h.UpdateContact).Methods(http.MethodPut)
	h.router.HandleFunc("/contacts/{id}", h.DeleteContact).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(serverPath, h.router))
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
		if errors.Is(err, db.ErrConflictRecord) {
			respondWithError(w, http.StatusConflict, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
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
