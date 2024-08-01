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

func NewBasicWaterLevelService(pinNumber int) (*BasicWaterLevelService, error) {
	err := rpio.Open()
	if err != nil {
		return nil, fmt.Errorf("Error opening GPIO: %w", err)
	}

	defer func() {
		if closeErr := rpio.Close(); closeErr != nil {
			log.Printf("error closing GPIO: %v", closeErr)
		}
	}()

	pin := rpio.Pin(pinNumber)
	pin.Input()

	return &BasicWaterLevelService{
		pin: pin,
	}, nil
}
