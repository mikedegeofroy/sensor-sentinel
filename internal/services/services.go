package services

type Services struct {
	WaterService WaterService
	AlarmService AlarmService
}

type WaterService interface {
	GetWaterLevel() (int, error)
	OnWaterLevelChange(func(int))
	StartPolling()
	StopPolling()
}

type AlarmService interface {
}
