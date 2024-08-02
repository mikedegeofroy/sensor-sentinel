package services

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type BasicWaterLevelService struct {
	pin        rpio.Pin
	callbacks  []func(int)
	polling    bool
	interval   time.Duration
	stopChan   chan struct{}
}

func initGPIO() error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("Error opening GPIO: %w", err)
	}
	return nil
}

func closeGPIO() {
	if err := rpio.Close(); err != nil {
		log.Printf("Error closing GPIO: %v", err)
	}
}

func (m *BasicWaterLevelService) GetWaterLevel() (int, error) {
	m.pin.PullUp()
	res := m.pin.Read()
	switch res {
	case rpio.Low:
		return 0, nil
	case rpio.High:
		return 1, nil
	default:
		return 0, fmt.Errorf("Unexpected GPIO pin reading")
	}
}

func (m *BasicWaterLevelService) OnWaterLevelChange(callback func(int)) {
	m.callbacks = append(m.callbacks, callback)
}

func NewBasicWaterLevelService(pinNumber int) (*BasicWaterLevelService, error) {
	if err := initGPIO(); err != nil {
		return nil, err
	}
	pin := rpio.Pin(pinNumber)
	pin.Input()

	return &BasicWaterLevelService{
		pin:      pin,
		interval: time.Second, // Default polling interval
		stopChan: make(chan struct{}),
	}, nil
}

func (m *BasicWaterLevelService) StartPolling() {
	m.polling = true
	go func() {
		var lastLevel int
		for {
			select {
			case <-m.stopChan:
				m.polling = false
				return
			default:
				level, err := m.GetWaterLevel()
				if err != nil {
					log.Printf("Error getting water level: %v", err)
					time.Sleep(m.interval)
					continue
				}
				if level != lastLevel {
					lastLevel = level
					for _, callback := range m.callbacks {
						callback(level)
					}
				}
				time.Sleep(m.interval)
			}
		}
	}()
}

func (m *BasicWaterLevelService) StopPolling() {
	close(m.stopChan)
	if m.polling {
		m.polling = false
	}
}
