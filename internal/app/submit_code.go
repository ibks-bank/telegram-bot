package app

import (
	"encoding/json"

	"github.com/ibks-bank/libs/cerr"
)

type submitCodeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type submitCodeResponse struct {
	Token string `json:"token"`
}

func (req *submitCodeRequest) Marshall() []byte {
	raw, _ := json.Marshal(req)
	return raw
}

func (a *app) parseSubmitCodeRequest(request []string) (*submitCodeRequest, error) {
	if len(request) != 3 {
		return nil, cerr.New("wrong number of args")
	}

	return &submitCodeRequest{
		Email:    request[0],
		Password: request[1],
		Code:     request[2],
	}, nil
}

func (a *app) submitCode(req *submitCodeRequest) (*submitCodeResponse, error) {
	respRaw, err := a.post(a.profileUrl+"/v1/auth/submit-code", "", req)
	if err != nil {
		return nil, cerr.Wrap(err, "can't do request")
	}

	return a.parseSubmitCodeResponse(respRaw)
}

func (a *app) parseSubmitCodeResponse(resp []byte) (*submitCodeResponse, error) {
	parsed := new(submitCodeResponse)
	err := json.Unmarshal(resp, parsed)
	if err != nil {
		return nil, cerr.Wrap(err, "can't unmarshall")
	}
	return parsed, nil
}
