package app

import (
	"encoding/json"
	"fmt"
	"github.com/ibks-bank/libs/cerr"
)

type getAccountRequest struct {
	AccountID string `json:"accountID"`
	Token     string `json:"-"`
}

type getAccountResponse struct {
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Limit    string `json:"limit"`
	Balance  string `json:"balance"`
}

func (resp *getAccountResponse) beautify() string {
	currency := formatCurrency(resp.Currency)

	return fmt.Sprintf(
		"Info about your bank account (%s):\n"+
			"Balance: %s %s\n"+
			"Limit: %s %s\n",
		resp.ID,
		resp.Balance, currency,
		resp.Limit, currency,
	)
}

func (a *app) parseGetAccountRequest(request []string) (*getAccountRequest, error) {
	if len(request) != 2 {
		return nil, cerr.New("wrong number of args")
	}

	return &getAccountRequest{
		AccountID: request[0],
		Token:     request[1],
	}, nil
}

func (a *app) getAccount(req *getAccountRequest) (*getAccountResponse, error) {
	endpoint := fmt.Sprintf("/v1/accounts/%s", req.AccountID)
	respRaw, err := get(a.bankAccountUrl+endpoint, req.Token)
	if err != nil {
		return nil, cerr.Wrap(err, "can't do request")
	}

	return a.parseGetAccountResponse(respRaw)
}

func (a *app) parseGetAccountResponse(resp []byte) (*getAccountResponse, error) {
	parsed := new(getAccountResponse)
	err := json.Unmarshal(resp, parsed)
	if err != nil {
		return nil, cerr.Wrap(err, "can't unmarshall")
	}
	return parsed, nil
}
