package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/denisboborukhin/contact_manager/internal/contact/manager"
	"github.com/denisboborukhin/contact_manager/internal/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "contact_manager"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	contactManager := manager.NewManager(db)
	contactHandler := handlers.NewContactHandler(contactManager)

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	handlers.SetupRoutes(router, contactHandler)

	port := getEnv("HTTP_PORT", "8080")
	log.Printf("Starting HTTP server on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
