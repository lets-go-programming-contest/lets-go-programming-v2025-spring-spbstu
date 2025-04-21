package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kseniadobrovolskaia/task-9/internal/phonebk"
)

type ContainerOfContacts interface {
	GetAll() ([]phonebk.Contact, error)
	Add(phonebk.Contact) error
	Update(phonebk.Contact) error
	Delete(string) error
	Get(string) (phonebk.Contact, error)
}

type Handler struct {
	pb ContainerOfContacts
}

func NewHandler(pb ContainerOfContacts) *Handler {
	return &Handler{pb: pb}
}

func (h *Handler) HandleContacts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		contacts, err := h.pb.GetAll()
		if err != nil {
			printError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)

	case http.MethodPost:
		var contact phonebk.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			printError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if contact.Name == "" || contact.Phone == "" {
			printError(w, "name and phone are required", http.StatusBadRequest)
			return
		}

		if err := h.pb.Add(contact); err != nil {
			printError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	default:
		printError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleOneContact(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/contacts/"):]
	if id == "" {
		printError(w, "ID is required", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		contact, err := h.pb.Get(id)
		if err != nil {
			printError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)

	case http.MethodPut:
		var contact phonebk.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			printError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if contact.Name == "" || contact.Phone == "" {
			printError(w, "name and new phone are required", http.StatusBadRequest)
			return
		}

		contact.ID = id
		if err := h.pb.Update(contact); err != nil {
			printError(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		err := h.pb.Delete(id)
		if err != nil {
			printError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	default:
		printError(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func printError(w http.ResponseWriter, msg string, status int) {
	http.Error(w, "{\"error_message\": \""+msg+"\"}", status)
}
