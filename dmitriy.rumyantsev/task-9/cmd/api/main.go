package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitriy.rumyantsev/task-9/internal/config"
	"github.com/dmitriy.rumyantsev/task-9/internal/handler"
	"github.com/dmitriy.rumyantsev/task-9/internal/repository"
	"github.com/dmitriy.rumyantsev/task-9/internal/repository/postgres"
	"github.com/dmitriy.rumyantsev/task-9/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	// Parse command line flags
	configPath := flag.String("config", "configs/app.yaml", "path to config file")
	flag.Parse()

	// Load configuration from YAML file
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("Starting application with configuration from %s", *configPath)

	// Connect to the database
	dbConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.Name,
	)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Configure database connection pool
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	// Check database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database")

	// Initialize repositories
	contactRepo := postgres.NewContactRepository(db)
	if err := contactRepo.InitSchema(); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	repos := repository.NewRepositories(contactRepo)

	// Initialize services
	deps := service.Dependencies{
		Repos: repos,
	}
	services := service.NewServices(deps)

	// Initialize HTTP handlers
	contactHandler := handler.NewContactHandler(services.Contact)

	// Initialize router
	mux := http.NewServeMux()
	contactHandler.Register(mux)

	// Configure HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      mux,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.Timeout * 2,
	}

	// Start server in a separate goroutine
	go func() {
		log.Printf("Starting server on port %s", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Configure graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout)
	defer cancel()

	// Stop the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
