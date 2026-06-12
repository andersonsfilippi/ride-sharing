package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/infrastructure/grpc"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/infrastructure/repository"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/service"
	grpcserver "google.golang.org/grpc"
)

var GrpcAddr = ":9093"

func main() {
	log.Println("Starting Trip Service")

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := new(sync.WaitGroup)

	wg.Go(func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	})

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Starting the grpc server
	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)

	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr().String())

	wg.Go(func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to server: %v", err)
			cancel()
		}
	})

	<-ctx.Done()
	grpcServer.GracefulStop()
	log.Println("Shutting down the server...")
	wg.Wait()
}
