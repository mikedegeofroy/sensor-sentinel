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

	ws.OnWaterLevelChange(func(int) {
		msg := tgbotapi.NewMessage(config.C.Telegram.ChatId, "penis")
		bot.Send(msg)
	})

	return &TelegramAlarmService{
		bot: bot,
	}, nil
}
