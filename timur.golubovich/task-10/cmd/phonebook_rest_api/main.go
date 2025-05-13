package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	gw "task-10/gen/proto/phonebook/v1"
	"task-10/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.ParseConfig("data/config.yaml")
	if err != nil {
		log.Fatalf("Fail to parse config file: %v", err)
	}

	grpcEndpoint := fmt.Sprintf("localhost:%d", cfg.GRPCPort)
	httpEndpoint := fmt.Sprintf("localhost:%d", cfg.HTTPPort)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := gw.RegisterPhonebookServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	log.Printf("HTTP server is running on port %s", httpEndpoint)
	if err := http.ListenAndServe(httpEndpoint, mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
