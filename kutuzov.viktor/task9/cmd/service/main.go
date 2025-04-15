package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	contactdatabase "github.com/vktr-ktzv/contact-api/internal/contactDatabase"
)

func main() {
	db := contactdatabase.ConnectDB()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/contacts", contactdatabase.GetContacts(db)).Methods("GET")
	router.HandleFunc("/contacts/{id}", contactdatabase.GetContactByID(db)).Methods("GET")
	router.HandleFunc("/contacts", contactdatabase.CreateContact(db)).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactdatabase.UpdateContact(db)).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contactdatabase.DeleteContact(db)).Methods("DELETE")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
