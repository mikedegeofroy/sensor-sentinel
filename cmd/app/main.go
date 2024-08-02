package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	server "sensor-sentinel/internal/gateways"
	services "sensor-sentinel/internal/services"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	svc := setupServices()

	s := server.NewServer(svc)
	go func() {
		if err := s.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("error during server shutdown: %v", err)
		}
	}()

	svc.WaterService.StartPolling()
	svc.WaterService.OnWaterLevelChange(func(level int) {
		fmt.Println("Level changed to %d", level)
	})

	<-ctx.Done()

	svc.WaterService.StopPolling()
	log.Println("Services stopped gracefully")
}

func setupServices() services.Services {
	waterService, err := services.NewBasicWaterLevelService(23)
	if err != nil {
		fmt.Println("Error initializing water service.")
		os.Exit(1)
	}

	return services.Services{
		WaterService: waterService,
	}
}
