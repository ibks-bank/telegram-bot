package app

import (
	"encoding/json"

	"github.com/ibks-bank/libs/cerr"
)

type payRequest struct {
	Token  string `json:"-"`
	Payee  string `json:"payee"`
	Amount string `json:"amount"`
}

func (req *payRequest) Marshall() []byte {
	raw, _ := json.Marshal(req)
	return raw
}

type payResponse struct {
}

func (resp *payResponse) beautify() string {
	return "Success!"
}

func (a *app) parsePayRequest(request []string) (*payRequest, error) {
	if len(request) != 3 {
		return nil, cerr.New("wrong number of args")
	}

	return &payRequest{
		Token:  request[2],
		Payee:  request[0],
		Amount: request[1],
	}, nil
}

func (a *app) pay(req *payRequest) (*payResponse, error) {
	respRaw, err := post(a.bankAccountUrl+"/v1/transactions/create", req.Token, req)
	if err != nil {
		return nil, cerr.Wrap(err, "can't do request")
	}

	return a.parsePayResponse(respRaw)
}

func (a *app) parsePayResponse(resp []byte) (*payResponse, error) {
	parsed := new(payResponse)
	err := json.Unmarshal(resp, parsed)
	if err != nil {
		return nil, cerr.Wrap(err, "can't unmarshall")
	}
	return parsed, nil
}
