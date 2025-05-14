package main

import (
	"fmt"
	"log"
	"net"

	v1 "task-10/gen/proto/phonebook/v1"
	"task-10/internal/config"
	phonebook "task-10/internal/handlers"
	"task-10/internal/manager"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.ParseConfig("data/config.yaml")
	if err != nil {
		log.Fatalf("Fail to parse config file: %v", err)
	}
	dsn := config.MakeString(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	pb := manager.NewManager(db)

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("Failed to create validator: %v", err)
	}

	handler := phonebook.NewHandler(pb, validator)

	grpcServer := grpc.NewServer()
	v1.RegisterPhonebookServiceServer(grpcServer, handler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %d", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
