package main

import (
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/ibks-bank/telegram-bot/config"
	"github.com/ibks-bank/telegram-bot/internal/app"
)

func main() {
	conf := config.GetConfig()

	bot, err := tgbotapi.NewBotAPI(conf.Auth.BotToken)
	if err != nil {
		log.Fatal("can't create bot")
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	app.New(
		bot,
		updateConfig,
		conf.Clients.ProfileURL,
		conf.Clients.BankAccountURL,
	).Run()
}
