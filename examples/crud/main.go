package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	addr := getenv("GRPC_ADDR", ":50051")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	s := grpc.NewServer(
		// For demo simplicity we use insecure creds. For TLS, set credentials here.
		grpc.Creds(insecure.NewCredentials()),
	// Add unary interceptors here if you want logging/metrics.
	)

	// Register services
	// TODO

	// Graceful shutdown handling
	go func() {
		log.Printf("gRPC server listening at %s", addr)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	// Give in-flight RPCs up to 5s to finish
	done := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("server stopped gracefully")
	case <-time.After(5 * time.Second):
		log.Println("force stop after timeout")
		s.Stop()
	}
}
