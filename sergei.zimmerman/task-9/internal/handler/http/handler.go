package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"phonebook/internal/service"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/contacts", h.GetAll).Methods("GET")
	r.HandleFunc("/contacts/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/contacts", h.Create).Methods("POST")
	r.HandleFunc("/contacts/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/contacts/{id}", h.Delete).Methods("DELETE")
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.svc.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch contacts", http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		http.Error(w, "Encoding response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.Parse(mux.Vars(r)["id"])

	c, err := h.svc.GetByID(id)
	if err != nil {
		http.Error(w, "Contact not found", http.StatusNotFound)

		return
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		http.Error(w, "Encoding response", http.StatusInternalServerError)
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
	}

	c, err := h.svc.Create(input.Name, input.Phone)
	if err != nil {
		http.Error(w, "Unable to create contact", http.StatusBadRequest)

		return
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		http.Error(w, "Encoding response", http.StatusInternalServerError)
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.Parse(mux.Vars(r)["id"])

	var input struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Encoding response", http.StatusInternalServerError)
	}

	err = h.svc.Update(id, input.Name, input.Phone)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.Parse(mux.Vars(r)["id"])

	err := h.svc.Delete(id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
