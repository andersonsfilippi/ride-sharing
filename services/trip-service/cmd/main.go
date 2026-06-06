package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	h "github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/infrastructure/http"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/infrastructure/repository"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/service"
)

func main() {

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)
	mux := http.NewServeMux()

	httpHandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	serverErr := make(chan error, 1)
	shutdown := make(chan os.Signal, 1)

	wg := new(sync.WaitGroup)

	wg.Go(func() {
		log.Printf("Server listening on %s", server.Addr)
		serverErr <- server.ListenAndServe()
	})

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		log.Printf("Error starting server: %v", err)
	case sig := <-shutdown:
		log.Printf("Server is shuttingdown due to %v signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not stop the server gracefully: %v", err)
			server.Close()
		}
	}
	wg.Wait()
}
