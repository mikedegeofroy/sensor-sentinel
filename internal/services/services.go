package services

type Services struct {
	WaterService WaterService
}

type WaterService interface {
	GetWaterLevel() (int, error)
}
