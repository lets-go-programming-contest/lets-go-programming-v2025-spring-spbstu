package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/quaiion/go-practice/contact-manager/internal/cm"
)

type ContactDB interface {
        GetAll() ([]cm.Contact, error)
	Get(string) (cm.Contact, error)
	Add(cm.Contact) error
	Update(cm.Contact) error
	Delete(string) error
}

type Handler struct {
	cdb ContactDB
}

func New(cdb ContactDB) *Handler {
	return &Handler{cdb: cdb}
}

func (h *Handler) HandleAllContacts(writer http.ResponseWriter, request *http.Request) {
        switch request.Method {
        case http.MethodGet:
                contacts, err := h.cdb.GetAll()
                if err != nil {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                json.NewEncoder(writer).Encode(contacts)

        case http.MethodPost:
                var contact cm.Contact

                err := json.NewDecoder(request.Body).Decode(&contact)
                if err != nil {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusBadRequest)
                        return
                }

                if contact.Name == `` || contact.Number == `` {
                        http.Error(writer, `{"error_message": "name and number required"}`, http.StatusBadRequest)
                        return
                }

                err = h.cdb.Add(contact)
                if err != nil {
                        if errors.Is(err, cm.ErrDuplicateAdded) {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusConflict)
                        } else {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusCreated)

        default:
                http.Error(writer, `{"error_message": "method not allowed"}`, http.StatusMethodNotAllowed)
                
        }
}

func (h *Handler) HandleContact(writer http.ResponseWriter, request *http.Request) {
        id := request.URL.Path[len(`/contacts/`):]
        if id == `` {
                http.Error(writer, `{"error_message": "ID required"}`, http.StatusBadRequest)
                return
        }

        switch request.Method {

        case http.MethodGet:
                contact, err := h.cdb.Get(id)
                if err != nil {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                json.NewEncoder(writer).Encode(contact)

        case http.MethodPut:
                var contact cm.Contact

                err := json.NewDecoder(request.Body).Decode(&contact)
                if err != nil {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusBadRequest)
                        return
                }

                if contact.Name == `` || contact.Number == `` {
                        http.Error(writer, `{"error_message": "name and new number required"}`, http.StatusBadRequest)
                        return
                }

                contact.ID = id
                err = h.cdb.Update(contact)
                if err != nil {
                        if errors.Is(err, cm.ErrContUpdNotFound) {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusNotFound)
                        } else if errors.Is(err, cm.ErrDuplicateAdded) {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusConflict)
                        } else {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusOK)

        case http.MethodDelete:
                err := h.cdb.Delete(id)
                if err != nil {
                        if errors.Is(err, cm.ErrContDelNotFound) {
                                http.Error(writer, `{"error_message": "contact not found"}`, http.StatusNotFound)
                        } else {
                                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                        }
                        return
                }

                writer.Header().Set("Content-Type", "application/json")
                writer.WriteHeader(http.StatusOK)

        default:
                http.Error(writer, `{"error_message": "method not allowed"}`, http.StatusMethodNotAllowed)
        }
}
