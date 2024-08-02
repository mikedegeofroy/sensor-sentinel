package services

import (
	"sensor-sentinel/cmd/app/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramAlarmService struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAlarmService(ws WaterService) (*TelegramAlarmService, error) {
	bot, err := tgbotapi.NewBotAPI(config.C.Telegram.ApiKey)
	if err != nil {
		return nil, err
	}

	ws.OnWaterLevelChange(func(level int) {
		if level == 0 {
			msg := tgbotapi.NewVenue(config.C.Telegram.ChatId, config.C.Telegram.Message, config.C.Cistern.Address, config.C.Cistern.Coordinates.Latitude, config.C.Cistern.Coordinates.Longitude)
			bot.Send(msg)
		}
	})

	return &TelegramAlarmService{
		bot: bot,
	}, nil
}
