package app

import (
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/ibks-bank/libs/cerr"
)

type app struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig

	profileUrl     string
	bankAccountUrl string
}

func New(
	bot *tgbotapi.BotAPI,
	updateConfig tgbotapi.UpdateConfig,
	profileUrl, bankAccountUrl string,
) *app {

	return &app{
		bot:            bot,
		updateConfig:   updateConfig,
		profileUrl:     profileUrl,
		bankAccountUrl: bankAccountUrl,
	}
}

func (a *app) Run() {
	updates, _ := a.bot.GetUpdatesChan(a.updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() != reflect.String || update.Message.Text == "" {
			continue
		}

		err := a.handle(update.Message)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, cerr.Wrap(err, "can't handle message").Error())
			a.bot.Send(msg)
		}
	}
}

func (a *app) handle(msg *tgbotapi.Message) error {
	splitted := strings.Split(msg.Text, " ")
	command := splitted[0]
	if command == msg.Text {
		return cerr.New("wrong args")
	}
	args := splitted[1:]

	var resp tgResponse

	switch command {
	case "/sign_in":

		req, err := a.parseSignInRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.signIn(req)
		if err != nil {
			return cerr.Wrap(err, "can't sign in")
		}

	case "/submit_code":

		req, err := a.parseSubmitCodeRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.submitCode(req)
		if err != nil {
			return cerr.Wrap(err, "can't submit code")
		}

	case "/passport":

		req, err := a.parseGetPassportRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.getPassport(req)
		if err != nil {
			return cerr.Wrap(err, "can't get passport")
		}

	case "/get_account":

		req, err := a.parseGetAccountRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.getAccount(req)
		if err != nil {
			return cerr.Wrap(err, "can't get account")
		}

	case "/get_accounts":

		req, err := a.parseGetAccountsRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.getAccounts(req)
		if err != nil {
			return cerr.Wrap(err, "can't get accounts")
		}

	case "/pay":

		req, err := a.parsePayRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.pay(req)
		if err != nil {
			return cerr.Wrap(err, "can't pay")
		}

	}

	a.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, resp.beautify()))

	return nil
}
