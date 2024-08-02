package config

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port        uint16 `envconfig:"HTTP_PORT" default:"8000"`
		AllowOrigin string `envconfig:"ALLOW_ORIGIN" default:"*"`
	}
	Telegram struct {
		ApiKey  string `envconfig:"TELEGRAM_API_KEY"`
		ChatId  int64  `envconfig:"TELEGRAM_CHAT_ID"`
		Message string `envconfig:"TELEGRAM_MESSAGE"`
	}
	Cistern struct {
		Coordinates struct {
			Latitude  float64 `envconfig:"CISTERN_LAT"`
			Longitude float64 `envconfig:"CISTERN_LON"`
		}
		Address string `envconfig:"CISTERN_ADDRESS"`
	}
}

var C Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("[INFO] no .env file, parsed exported variables")
	}
	err = envconfig.Process("", &C)
	if err != nil {
		log.Fatalf("can't parse config: %s", err)
	}

	printConfig(C)
}

func printConfig(c Config) {
	data, _ := json.MarshalIndent(c, "", "\t")
	fmt.Println("=== CONFIG ===")
	fmt.Println(string(data))
	fmt.Println("==============")
}
