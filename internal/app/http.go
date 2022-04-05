package app

import (
	"bytes"
	"io"
	"net/http"
)

type body interface {
	Marshall() []byte
}

type tgResponse interface {
	beautify() string
}

func get(url, token string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return send(prepareHeaders(req, "Content-Type", "application/json", "x-auth-token", token))
}

func post(url, token string, body body) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body.Marshall()))
	if err != nil {
		return nil, err
	}

	return send(prepareHeaders(req, "Content-Type", "application/json", "x-auth-token", token))
}

func send(req *http.Request) ([]byte, error) {

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func prepareHeaders(req *http.Request, kvHeaders ...string) *http.Request {
	if len(kvHeaders)%2 == 1 {
		return req
	}

	for i := 0; i < len(kvHeaders)-2; i += 2 {
		req.Header.Set(kvHeaders[i], kvHeaders[i+1])
	}

	return req
}
