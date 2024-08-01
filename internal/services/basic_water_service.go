package services

import (
	"fmt"
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

type BasicWaterLevelService struct {
	pin rpio.Pin
}

// Initialize GPIO once at the start of your application
func initGPIO() error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("Error opening GPIO: %w", err)
	}
	return nil
}

// Close GPIO once at the end of your application
func closeGPIO() {
	if err := rpio.Close(); err != nil {
		log.Printf("Error closing GPIO: %v", err)
	}
}

// Implement the GetWaterLevel method with error handling
func (m *BasicWaterLevelService) GetWaterLevel() (int, error) {
	initGPIO()
	res := m.pin.Read()
	closeGPIO()
	switch res {
	case rpio.Low:
		return 0, nil
	case rpio.High:
		return 1, nil
	default:
		return 0, fmt.Errorf("Unexpected GPIO pin reading")
	}
}

// Initialize and return the BasicWaterLevelService
func NewBasicWaterLevelService(pinNumber int) (*BasicWaterLevelService, error) {
	pin := rpio.Pin(pinNumber)
	pin.Input()

	return &BasicWaterLevelService{
		pin: pin,
	}, nil
}
