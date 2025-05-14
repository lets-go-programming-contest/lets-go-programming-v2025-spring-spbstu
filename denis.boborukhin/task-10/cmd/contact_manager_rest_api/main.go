package main

import (
	"context"
	"log"
	"net/http"

	gw "github.com/denisboborukhin/contact_manager/gen/proto/contact_manager/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcEndpoint := "localhost:50051"
	httpEndpoint := "localhost:8080"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := gw.RegisterContactManagerServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	log.Printf("HTTP server is running on port %s", httpEndpoint)
	if err := http.ListenAndServe(httpEndpoint, mux); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
