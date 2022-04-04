package app

import (
	"encoding/json"

	"github.com/ibks-bank/libs/cerr"
)

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *signInRequest) Marshall() []byte {
	raw, _ := json.Marshal(req)
	return raw
}

func (a *app) parseSignInRequest(request []string) (*signInRequest, error) {
	if len(request) != 2 {
		return nil, cerr.New("wrong number of args")
	}

	return &signInRequest{Email: request[0], Password: request[1]}, nil
}

func (a *app) signIn(req *signInRequest) error {
	_, err := a.post(a.bankAccountUrl+"/v1/auth/sign-in", "", req)
	return cerr.Wrap(err, "can't do request")
}
