package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quaiion/go-practice/contact-manager/internal/cm"
)

type ContactDB interface {
        GetAll() ([]cm.Contact, error)
	Get(string) (cm.Contact, error)
	Add(cm.Contact) error
	Update(cm.Contact) error
	Delete(string) error
}

type handler struct {
	cdb ContactDB
}

func New(rout *mux.Router, cdb ContactDB) *mux.Router {
        hand := &handler{ cdb: cdb }

        rout.HandleFunc("/contacts/{id}", hand.get   ).Methods(http.MethodOptions, http.MethodGet)
        rout.HandleFunc("/contacts",      hand.getAll).Methods(http.MethodOptions, http.MethodGet)
        rout.HandleFunc("/contacts",      hand.add   ).Methods(http.MethodOptions, http.MethodPost)
        rout.HandleFunc("/contacts/{id}", hand.update).Methods(http.MethodOptions, http.MethodPut)
        rout.HandleFunc("/contacts/{id}", hand.delete).Methods(http.MethodOptions, http.MethodDelete)

        return rout
}

func (hand *handler) get(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars["id"]
        if id == `` {
                http.Error(writer, `{"error_message": "ID required"}`, http.StatusBadRequest)
                return
        }

        contact, err := hand.cdb.Get(id)
        if err != nil {
                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                return
        }

        writer.Header().Set("Content-Type", "application/json")
        json.NewEncoder(writer).Encode(contact)
}

func (hand *handler) getAll(writer http.ResponseWriter, request *http.Request) {
        contacts, err := hand.cdb.GetAll()
        if err != nil {
                http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                return
        }

        writer.Header().Set("Content-Type", "application/json")
        json.NewEncoder(writer).Encode(contacts)
}

func (hand *handler) add(writer http.ResponseWriter, request *http.Request) {
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

        err = hand.cdb.Add(contact)
        if err != nil {
                if errors.Is(err, cm.ErrDuplicateAdded) {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusConflict)
                } else {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                }
                return
        }

        writer.Header().Set("Content-Type", "application/json")
}

func (hand *handler) update(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars["id"]
        if id == `` {
                http.Error(writer, `{"error_message": "ID required"}`, http.StatusBadRequest)
                return
        }

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
        err = hand.cdb.Update(contact)
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
}

func (hand *handler) delete(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars["id"]
        if id == `` {
                http.Error(writer, `{"error_message": "ID required"}`, http.StatusBadRequest)
                return
        }

        err := hand.cdb.Delete(id)
        if err != nil {
                if errors.Is(err, cm.ErrContDelNotFound) {
                        http.Error(writer, `{"error_message": "contact not found"}`, http.StatusNotFound)
                } else {
                        http.Error(writer, `{"error_message": "` + err.Error() + `"}`, http.StatusInternalServerError)
                }
                return
        }

        writer.Header().Set("Content-Type", "application/json")
}
