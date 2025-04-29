package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/denisboborukhin/contact_manager/gen/proto/contact_manager/v1"
	"github.com/denisboborukhin/contact_manager/internal/contact/manager"
	"github.com/denisboborukhin/contact_manager/internal/handlers/contact_manager"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
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

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("Failed to create validator: %v", err)
	}

	handler := contact_manager.NewHandler(contactManager, validator)

	grpcServer := grpc.NewServer()
	v1.RegisterContactManagerServiceServer(grpcServer, handler)

	grpcPort := getEnv("GRPC_PORT", "50051")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
