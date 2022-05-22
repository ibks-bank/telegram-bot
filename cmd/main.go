package main

import (
	"database/sql"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/ibks-bank/telegram-bot/config"
	"github.com/ibks-bank/telegram-bot/internal/app"
	"github.com/ibks-bank/telegram-bot/internal/store"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.GetConfig()

	bot, err := tgbotapi.NewBotAPI(conf.Auth.BotToken)
	if err != nil {
		log.Fatal("can't create bot")
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	postgres, err := sql.Open(
		"postgres",
		"port=5433 host=localhost user=postgres password=postgres dbname=telegram sslmode=disable",
	)
	if err != nil {
		log.Fatal("can't open database")
	}

	app.New(
		bot,
		updateConfig,
		store.New(postgres),
		conf.Clients.ProfileURL,
		conf.Clients.BankAccountURL,
	).Run()
}
