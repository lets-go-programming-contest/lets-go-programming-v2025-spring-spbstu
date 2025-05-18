package handlers

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, handler *ContactHandler) {
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/contacts", handler.GetContacts).Methods("GET")
	v1.HandleFunc("/contacts/{id}", handler.GetContact).Methods("GET")
	v1.HandleFunc("/contacts", handler.CreateContact).Methods("POST")
	v1.HandleFunc("/contacts/{id}", handler.UpdateContact).Methods("PUT")
	v1.HandleFunc("/contacts/{id}", handler.DeleteContact).Methods("DELETE")
}
