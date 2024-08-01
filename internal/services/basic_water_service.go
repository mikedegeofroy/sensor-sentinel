package services

import (
	"fmt"
	"log"
	"github.com/stianeikeland/go-rpio/v4"
)

type BasicWaterLevelService struct {
	pin rpio.Pin
}

// Implement the GetWaterLevel method with error handling
func (m *BasicWaterLevelService) GetWaterLevel() (int, error) {
	res := m.pin.Read()
	if res == rpio.Low {
		return 0, nil
	} else if res == rpio.High {
		return 1, nil
	} else {
		return 0, fmt.Errorf("Unexpected GPIO pin reading")
	}
}

// Initialize and return the BasicWaterLevelService
func NewBasicWaterLevelService(pinNumber int) (*BasicWaterLevelService, error) {
	// Open GPIO
	if err := rpio.Open(); err != nil {
		return nil, fmt.Errorf("Error opening GPIO: %w", err)
	}

	// Initialize pin
	pin := rpio.Pin(pinNumber)
	pin.Input()

	return &BasicWaterLevelService{
		pin: pin,
	}, nil
}

// Close GPIO when the service is no longer needed
func (m *BasicWaterLevelService) Close() {
	if err := rpio.Close(); err != nil {
		log.Printf("Error closing GPIO: %v", err)
	}
}
