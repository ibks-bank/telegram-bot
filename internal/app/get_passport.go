package app

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ibks-bank/libs/cerr"
)

type getPassportRequest struct {
	Token string `json:"-"`
}

type getPassportResponse struct {
	ID         int64     `json:"id"`
	Series     string    `json:"series"`
	Number     string    `json:"number"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	IssuedBy   string    `json:"issued_by"`
	IssuedAt   time.Time `json:"issued_at"`
	Address    string    `json:"address"`
	Birthplace string    `json:"birthplace"`
	Birthdate  time.Time `json:"birthdate"`
}

func (resp *getPassportResponse) beautify() string {
	return fmt.Sprintf(
		"Your passport is:\n"+
			"Series: %s\n"+
			"Number: %s\n"+
			"First name: %s\n"+
			"Middle name: %s\n"+
			"Last name: %s\n"+
			"Issued by: %s\n"+
			"Issued at: %s\n"+
			"Address: %s\n"+
			"Birthplace: %s\n"+
			"Birhdate: %s",
		resp.Series, resp.Number, resp.FirstName, resp.MiddleName, resp.LastName,
		resp.IssuedBy, formatDate(resp.IssuedAt), resp.Address, resp.Birthplace, formatDate(resp.Birthdate),
	)
}

func (a *app) parseGetPassportRequest(request []string) (*getPassportRequest, error) {
	if len(request) != 1 {
		return nil, cerr.New("wrong number of args")
	}

	return &getPassportRequest{Token: request[0]}, nil
}

func (a *app) getPassport(req *getPassportRequest) (*getPassportResponse, error) {
	respRaw, err := get(a.profileUrl+"/v1/passport", req.Token)
	if err != nil {
		return nil, cerr.Wrap(err, "can't do request")
	}

	return a.parseGetPassportResponse(respRaw)
}

func (a *app) parseGetPassportResponse(resp []byte) (*getPassportResponse, error) {
	parsed := new(getPassportResponse)
	err := json.Unmarshal(resp, parsed)
	if err != nil {
		return nil, cerr.Wrap(err, "can't unmarshall")
	}
	return parsed, nil
}
