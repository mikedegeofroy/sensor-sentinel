package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	server "sensor-sentinel/internal/gateways"
)

// @title           SensorSentinel
// @version         2.0
// @description     Notify when water level is low

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := server.NewServer()
	if err := s.Run(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("error during server shutdown: %v", err)
	}
}
