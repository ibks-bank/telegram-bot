package app

import (
	"encoding/json"

	"github.com/ibks-bank/libs/cerr"
)

type getAccountsRequest struct {
	Token string `json:"-"`
}

type getAccountsResponse struct {
	Accounts []*getAccountResponse `json:"accounts"`
}

func (resp *getAccountsResponse) beautify() string {
	result := ""
	for _, account := range resp.Accounts {
		result += account.beautify()
	}
	return result
}

func (a *app) parseGetAccountsRequest(request []string) (*getAccountsRequest, error) {
	if len(request) != 1 {
		return nil, cerr.New("wrong number of args")
	}

	return &getAccountsRequest{
		Token: request[0],
	}, nil
}

func (a *app) getAccounts(req *getAccountsRequest) (*getAccountsResponse, error) {
	respRaw, err := get(a.bankAccountUrl+"/v1/accounts", req.Token)
	if err != nil {
		return nil, cerr.Wrap(err, "can't do request")
	}

	return a.parseGetAccountsResponse(respRaw)
}

func (a *app) parseGetAccountsResponse(resp []byte) (*getAccountsResponse, error) {
	parsed := new(getAccountsResponse)
	err := json.Unmarshal(resp, parsed)
	if err != nil {
		return nil, cerr.Wrap(err, "can't unmarshall")
	}
	return parsed, nil
}
