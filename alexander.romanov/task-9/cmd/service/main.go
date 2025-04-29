package main

import (
	"github.com/alexander-romanov-edu/phonebook/internal/handler"
	"github.com/alexander-romanov-edu/phonebook/internal/repository"
	"github.com/alexander-romanov-edu/phonebook/internal/service"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewPhonebookRepository()
	defer repo.DeleteMe()

	// Initialize database
	if err := repo.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	svc := service.NewPhonebookService(*repo)
	h := handler.NewPhonebookHandler(*svc)

	http.HandleFunc("/contacts", h.HandleContacts)
	http.HandleFunc("/contacts/", h.HandleSingleContact)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
