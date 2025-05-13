package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	contactpb "task-10/gen/proto/api/proto/contact_manager/v1"
	"task-10/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	grpcAddr := fmt.Sprintf("localhost:%s", cfg.GRPC.Port)
	restAddr := fmt.Sprintf(":%s", cfg.REST.Port)

	conn, err := net.DialTimeout("tcp", grpcAddr, 3*time.Second)
	if err != nil {
		log.Fatalf("Cannot connect to gRPC server: %v", err)
	}
	conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := contactpb.RegisterContactServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}

	server := &http.Server{
		Addr:    restAddr,
		Handler: mux,
	}

	go func() {
		log.Printf("REST server running on %s", restAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	log.Println("Shutting down server...")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
}
