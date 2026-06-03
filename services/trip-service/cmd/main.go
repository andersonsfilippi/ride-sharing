package main

import (
	"context"
	"log"
	"time"

	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/domain"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/infrastructure/repository"
	"github.com/andersonsfilippi/ride-sharing/services/trip-service/internal/service"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInmemRepository()

	fare := &domain.RideFareModel{
		UserID: "42",
	}

	svc := service.NewService(inmemRepo)

	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}
	log.Println(t)

	// keep the program running for now
	for {
		time.Sleep(time.Second)
	}
}
