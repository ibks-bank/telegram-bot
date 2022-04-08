package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	Auth    *AuthConfiguration
	Clients *Clients
}

type AuthConfiguration struct {
	BotToken string
}

type Clients struct {
	ProfileURL     string
	BankAccountURL string
}

var config *configuration

func GetConfig() *configuration {
	if config == nil {
		config = readConfig()
	}

	return config
}

func readConfig() *configuration {
	err := godotenv.Load("./config/dev.env")
	if err != nil {
		log.Println("Can't load config file")
	}

	return &configuration{
		Auth: &AuthConfiguration{
			BotToken: os.Getenv("BOT_TOKEN"),
		},
		Clients: &Clients{
			ProfileURL:     os.Getenv("CLIENTS_PROFILE_URL"),
			BankAccountURL: os.Getenv("CLIENTS_BANK_ACCOUNT_URL"),
		},
	}
}
