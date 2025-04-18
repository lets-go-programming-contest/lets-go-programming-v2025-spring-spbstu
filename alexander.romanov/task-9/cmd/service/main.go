package main

import (
	"github.com/alexander-romanov-edu/phonebook/internal/handler"
	"github.com/alexander-romanov-edu/phonebook/internal/repository"
	"github.com/alexander-romanov-edu/phonebook/internal/service"
	"log"
	"net/http"
	"os"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "phonebook")

	repo := repository.NewPhonebookRepository(dbHost, dbPort, dbUser, dbPassword, dbName)
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
