package main

import (
	"log"
	"net/http"

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

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
