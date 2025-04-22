package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task-9/internal/api"
	"task-9/internal/config"
	"task-9/internal/storage"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := storage.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	contactStore := storage.NewContactStore(db)
	if err := contactStore.Initialize(); err != nil {
		log.Fatalf("Failed to initialize contact store: %v", err)
	}

	handler := api.NewHandler(contactStore)

	//See: https://pkg.go.dev/net/http
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/contacts", handler.HandleContacts)
	http.HandleFunc("/contacts/", handler.HandleContact)

	go func() {
		log.Printf("Server started on port %s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")
}
