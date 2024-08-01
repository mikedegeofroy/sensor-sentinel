package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	server "sensor-sentinel/internal/gateways"
	services "sensor-sentinel/internal/services"
)

// @title           SensorSentinel
// @version         2.0
// @description     Notify when water level is low

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := server.NewServer(setupServices())
	if err := s.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("error during server shutdown: %v", err)
	}
}

func setupServices() services.Services {
	waterService, err := services.NewBasicWaterLevelService(23)
	if err != nil {
		fmt.Println("Error initializing water service.")
	}

	return services.Services{
		WaterService: waterService,
	}
}
