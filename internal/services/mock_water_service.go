package services

import (
	"time"
)

type MockWaterLevelService struct {
	callbacks []func(int)
	polling   bool
	interval  time.Duration
	stopChan  chan struct{}
}

func (m *MockWaterLevelService) GetWaterLevel() (int, error) {
	return 10, nil
}

func (m *MockWaterLevelService) OnWaterLevelChange(callback func(int)) {
	m.callbacks = append(m.callbacks, callback)
}

func NewMockWaterLevelService(pinNumber int) (*MockWaterLevelService, error) {

	return &MockWaterLevelService{
		interval: time.Second,
		stopChan: make(chan struct{}),
	}, nil
}

func (m *MockWaterLevelService) StartPolling() {
	for _, callback := range m.callbacks {
		callback(0)
	}
}

func (m *MockWaterLevelService) StopPolling() {
}
