package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quaiion/go-practice/contact-manager/internal/cm"
	"github.com/quaiion/go-practice/contact-manager/internal/httperr"
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

var (
        errDecodeFailed     = errors.New("failed decoding the input contact")
        errEmptyID          = errors.New("input ID should not be empty")
        errEmptyName        = errors.New("input name should not be empty")
)

func (hand *handler) get(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars[`id`]
        if id == `` {
                httperr.New(errEmptyID).Encode(writer, http.StatusBadRequest)
                return
        }

        contact, err := hand.cdb.Get(id)
        if err != nil {
                if errors.Is(err, cm.ErrGetContNotFound) {
                        httperr.New(err).Encode(writer, http.StatusNotFound)
                } else {
                        httperr.New(err).Encode(writer, http.StatusInternalServerError)
                }
                return
        }

        setHeader(writer)
        json.NewEncoder(writer).Encode(contact)
}

func (hand *handler) getAll(writer http.ResponseWriter, request *http.Request) {
        contacts, err := hand.cdb.GetAll()
        if err != nil {
                httperr.New(err).Encode(writer, http.StatusInternalServerError)
                return
        }

        setHeader(writer)
        json.NewEncoder(writer).Encode(contacts)
}

func (hand *handler) add(writer http.ResponseWriter, request *http.Request) {
        var contact cm.Contact

        err := json.NewDecoder(request.Body).Decode(&contact)
        if err != nil {
                httperr.New(errors.Join(errDecodeFailed, err)).Encode(writer, http.StatusBadRequest)
                return
        }

        if contact.Name == `` {
                httperr.New(errEmptyName).Encode(writer, http.StatusBadRequest)
                return
        }

        err = hand.cdb.Add(contact)
        if err != nil {
                if errors.Is(err, cm.ErrDuplicateAdded) {
                        httperr.New(err).Encode(writer, http.StatusConflict)
                } else {
                        httperr.New(err).Encode(writer, http.StatusInternalServerError)
                }
                return
        }

        setHeader(writer)
}

func (hand *handler) update(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars[`id`]
        if id == `` {
                httperr.New(errEmptyID).Encode(writer, http.StatusBadRequest)
                return
        }

        var contact cm.Contact

        err := json.NewDecoder(request.Body).Decode(&contact)
        if err != nil {
                httperr.New(errors.Join(errDecodeFailed, err)).Encode(writer, http.StatusBadRequest)
                return
        }

        if contact.Name == `` {
                httperr.New(errEmptyName).Encode(writer, http.StatusBadRequest)
                return
        }

        contact.ID = id
        err = hand.cdb.Update(contact)
        if err != nil {
                if errors.Is(err, cm.ErrContUpdNotFound) {
                        httperr.New(err).Encode(writer, http.StatusNotFound)
                } else if errors.Is(err, cm.ErrDuplicateAdded) {
                        httperr.New(err).Encode(writer, http.StatusConflict)
                } else {
                        httperr.New(err).Encode(writer, http.StatusInternalServerError)
                }
                return
        }

        setHeader(writer)
}

func (hand *handler) delete(writer http.ResponseWriter, request *http.Request) {
        vars := mux.Vars(request)
	id := vars[`id`]
        if id == `` {
                httperr.New(errEmptyID).Encode(writer, http.StatusBadRequest)
                return
        }

        err := hand.cdb.Delete(id)
        if err != nil {
                if errors.Is(err, cm.ErrContDelNotFound) {
                        httperr.New(err).Encode(writer, http.StatusNotFound)
                } else {
                        httperr.New(err).Encode(writer, http.StatusInternalServerError)
                }
                return
        }

        setHeader(writer)
}

func setHeader(writer http.ResponseWriter) {
        writer.Header().Set("Content-Type", "application/json")
}
