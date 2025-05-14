package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	contactpb "task-10/gen/proto/api/proto/contact_manager/v1"
	"task-10/internal/config"
	"task-10/internal/handler"
	"task-10/internal/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcServer := grpc.NewServer()
	contactService := handler.NewContactServiceHandler(contactStore)
	contactpb.RegisterContactServiceServer(grpcServer, contactService)
	reflection.Register(grpcServer)

	grpcAddr := fmt.Sprintf(":%s", cfg.GRPC.Port)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Starting gRPC server on port %s", cfg.GRPC.Port)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
}
