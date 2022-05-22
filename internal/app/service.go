package app

import (
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/ibks-bank/libs/cerr"
)

type store interface {
	InsertUser(username string) error
	UpdateToken(username, token string) error
	GetToken(username string) (string, error)
}

type app struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	store        store

	profileUrl     string
	bankAccountUrl string
}

func New(
	bot *tgbotapi.BotAPI,
	updateConfig tgbotapi.UpdateConfig,
	store store,
	profileUrl, bankAccountUrl string,
) *app {

	return &app{
		bot:            bot,
		updateConfig:   updateConfig,
		store:          store,
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
	//if command == msg.Text {
	//	return cerr.New("wrong args")
	//}
	args := make([]string, 0)
	if len(splitted) > 1 {
		args = append(args, splitted[1:]...)
	}

	var resp tgResponse

	err := a.store.InsertUser(msg.From.UserName)
	if err != nil {
		return cerr.Wrap(err, "can't insert user")
	}

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

	case "/code":

		req, err := a.parseSubmitCodeRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		resp, err = a.submitCode(req)
		if err != nil {
			return cerr.Wrap(err, "can't submit code")
		}

		err = a.store.UpdateToken(msg.From.UserName, resp.beautify())
		if err != nil {
			return cerr.Wrap(err, "can't update token")
		}

		resp = &submitCodeResponse{Token: "Successfully logged in!"}

	case "/passport":

		//req, err := a.parseGetPassportRequest(args)
		//if err != nil {
		//	return cerr.Wrap(err, "can't parse request")
		//}

		token, err := a.store.GetToken(msg.From.UserName)
		if err != nil {
			return cerr.Wrap(err, "can't get token")
		}

		resp, err = a.getPassport(&getPassportRequest{Token: token})
		if err != nil {
			return cerr.Wrap(err, "can't get passport")
		}

	case "/account":

		req, err := a.parseGetAccountRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		token, err := a.store.GetToken(msg.From.UserName)
		if err != nil {
			return cerr.Wrap(err, "can't get token")
		}
		req.Token = token

		resp, err = a.getAccount(req)
		if err != nil {
			return cerr.Wrap(err, "can't get account")
		}

	case "/accounts":

		//req, err := a.parseGetAccountsRequest(args)
		//if err != nil {
		//	return cerr.Wrap(err, "can't parse request")
		//}

		token, err := a.store.GetToken(msg.From.UserName)
		if err != nil {
			return cerr.Wrap(err, "can't get token")
		}

		resp, err = a.getAccounts(&getAccountsRequest{Token: token})
		if err != nil {
			return cerr.Wrap(err, "can't get accounts")
		}

	case "/pay":

		req, err := a.parsePayRequest(args)
		if err != nil {
			return cerr.Wrap(err, "can't parse request")
		}

		token, err := a.store.GetToken(msg.From.UserName)
		if err != nil {
			return cerr.Wrap(err, "can't get token")
		}
		req.Token = token

		resp, err = a.pay(req)
		if err != nil {
			return cerr.Wrap(err, "can't pay")
		}

	default:

		resp = &defaultResp{text: "wrong command"}

	}

	a.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, resp.beautify()))

	return nil
}
