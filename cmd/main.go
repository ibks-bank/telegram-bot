package main

import (
	"context"
	"log"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/ibks-bank/telegram-bot/internal/app"
)

func main() {
	ctx := context.Background()

	bot, err := tgbotapi.NewBotAPI("5128594721:AAE8g9OWzW2W7yRGZvJzZ755O1Vz8_gDZLk")
	if err != nil {
		log.Fatal("can't create bot")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	a := app.New()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() != reflect.String || update.Message.Text == "" {
			continue
		}

		switch update.Message.Text {
		case "/sign_in":
			resp, err := a.SignIn(ctx, &app.SignInRequest{})
			if err != nil {
				panic(err)
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Token)
			bot.Send(msg)
		}
	}
}
